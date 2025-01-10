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
			// 用户管理 system:user:xxx
			userGroup := authorized.Group("/system/user")
			{
				userGroup.GET("", handler.User().List)                      // system:user:list
				userGroup.POST("", handler.User().Create)                   // system:user:create
				userGroup.PUT("/:id", handler.User().Update)                // system:user:update
				userGroup.DELETE("/:id", handler.User().Delete)             // system:user:delete
				userGroup.GET("/:id", handler.User().Detail)                // system:user:detail
				userGroup.PATCH("/password", handler.User().UpdatePassword) // system:user:password
				userGroup.POST("/assign", handler.User().AssignRoles)       // system:user:assign
			}

			// 角色管理 system:role:xxx
			roleGroup := authorized.Group("/system/role")
			{
				roleGroup.GET("", handler.Role().List)                // system:role:list
				roleGroup.POST("", handler.Role().Create)             // system:role:create
				roleGroup.PUT("/:id", handler.Role().Update)          // system:role:update
				roleGroup.DELETE("/:id", handler.Role().Delete)       // system:role:delete
				roleGroup.GET("/:id", handler.Role().Detail)          // system:role:detail
				roleGroup.GET("/menus", handler.Role().GetMenus)      // system:role:menus
				roleGroup.POST("/assign", handler.Role().AssignMenus) // system:role:assign
			}

			// 菜单管理 system:menu:xxx
			menuGroup := authorized.Group("/system/menu")
			{
				//menuGroup.GET("", handler.Menu().List)             // system:menu:list
				menuGroup.POST("", handler.Menu().Create)          // system:menu:create
				menuGroup.PUT("/:id", handler.Menu().Update)       // system:menu:update
				menuGroup.DELETE("/:id", handler.Menu().Delete)    // system:menu:delete
				menuGroup.GET("/tree", handler.Menu().GetMenuTree) // system:menu:tree
			}
		}
	}

	return r
}
