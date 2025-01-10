package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/nunu-casbin-admin/internal/handler"
	"github.com/wxlbd/nunu-casbin-admin/internal/middleware"
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
	"github.com/wxlbd/nunu-casbin-admin/pkg/config"
	"github.com/wxlbd/nunu-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/nunu-casbin-admin/pkg/log"
)

func NewServerHTTP(
	cfg *config.Config,
	logger *log.Logger,
	jwt *jwtx.JWT,
	handler *handler.Handler,
	enforcer *casbin.Enforcer,
	svc service.Service,
) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// 管理后台接口
	admin := r.Group("/admin/v1")
	{
		// 完全公开的接口
		admin.POST("/login", handler.User().Login)
		admin.POST("/refresh-token", handler.User().RefreshToken)

		// 需要JWT认证的接口
		jwtAuth := admin.Group("").Use(middleware.JWTAuth(jwt))
		{
			jwtAuth.GET("/user/current", handler.User().Current)
			jwtAuth.GET("/me/menus", handler.Menu().GetUserMenus)
			jwtAuth.GET("/menu/tree", handler.Menu().GetMenuTree)
			jwtAuth.GET("user/roles", handler.User().GetRoles)
		}

		// 需要完整权限控制的接口
		authorized := admin.Group("")
		authorized.Use(
			middleware.JWTAuth(jwt),
			middleware.CasbinMiddleware(enforcer, logger, svc),
		)
		{
			// 用户管理
			user := authorized.Group("/user")
			{
				user.GET("", handler.User().List)
				user.POST("", handler.User().Create)
				user.PUT("/:id", handler.User().Update)
				user.DELETE("/:id", handler.User().Delete)
				user.POST("/password", handler.User().UpdatePassword)
				user.POST("/roles", handler.User().AssignRoles)
			}

			// 角色管理
			role := authorized.Group("/role")
			{
				role.GET("", handler.Role().List)
				role.POST("", handler.Role().Create)
				role.PUT("/:id", handler.Role().Update)
				role.DELETE("/:id", handler.Role().Delete)
				role.GET("/menus", handler.Role().GetMenus)
				role.POST("/menus", handler.Role().AssignMenus)
			}

			// 菜单管理
			menu := authorized.Group("/menu")
			{
				menu.POST("", handler.Menu().Create)
				menu.PUT("/:id", handler.Menu().Update)
				menu.DELETE("/:id", handler.Menu().Delete)
			}
		}
	}

	return r
}
