package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wxlbd/gin-casbin-admin/internal/handler"
	"github.com/wxlbd/gin-casbin-admin/internal/middleware"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
)

func NewServerHTTP(
	cfg *config.Config,
	logger *log.Logger,
	jwt *jwtx.JWT,
	handler *handler.Handler,
	enforcer *casbin.Enforcer,
	svc handler.Service,
) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
		middleware.RequestLogger(logger),
		middleware.ErrorHandler(),
	)
	api := r.Group("api")
	{
		auth := api.Group("auth")
		// 完全公开的接口
		auth.POST("/login", handler.User().Login)
		auth.POST("/refresh-token", handler.User().RefreshToken)
		auth.POST("/logout", handler.User().Logout)
		auth.GET("/captcha", handler.Captcha().Generate)

		// 需要JWT认证的接口
		user := api.Group("user")
		user.Use(middleware.JWTAuth(jwt))
		{
			profile := user.Group("profile")
			{
				profile.GET("", handler.User().Current)
				profile.GET("menus", handler.SysMenu().GetUserMenuTree)
				// profile.GET("/menu/tree", handler.Menu().GetMenuTree)
				profile.GET("roles", handler.User().GetCurrentUserRoles)
			}
		}

		// 需要完整权限控制的接口
		authorized := api.Group("")
		authorized.Use(
			middleware.JWTAuth(jwt),
			middleware.CasbinMiddleware(enforcer, logger, svc),
		)
		sys := authorized.Group("system")

		// 权限控制
		{
			// 用户管理 system:user:xxx
			userGroup := sys.Group("user")
			{
				userGroup.GET("", handler.User().List)                         // permission:user:list
				userGroup.POST("", handler.User().Create)                      // permission:user:create
				userGroup.PUT("/:id", handler.User().Update)                   // permission:user:update
				userGroup.DELETE("/:ids", handler.User().Delete)               // permission:user:delete
				userGroup.GET("/:id", handler.User().Detail)                   // permission:user:detail
				userGroup.GET("/:id/roles", handler.User().GerUserRoles)       // permission:user:get:roles
				userGroup.PATCH(":id/password", handler.User().UpdatePassword) // permission:user:set:password
				userGroup.PUT(":id/roles", handler.User().AssignRoles)         // permission:user:set:roles
			}

			// 角色管理 permission:role:xxx
			roleGroup := sys.Group("role")
			{
				roleGroup.GET("", handler.Role().List)                           // permission:role:list
				roleGroup.POST("", handler.Role().Create)                        // permission:role:create
				roleGroup.PUT("/:id", handler.Role().Update)                     // permission:role:update
				roleGroup.DELETE("/:ids", handler.Role().Delete)                 // permission:role:delete
				roleGroup.GET("/:id", handler.Role().Detail)                     // permission:role:detail
				roleGroup.GET("/:id/menus", handler.Role().GetPermittedMenus)    // permission:role:get:menus
				roleGroup.PUT("/:id/menus", handler.Role().AssignRoleMenusByIDs) // permission:role:set:menus
			}

			// 菜单管理 permission:menu:xxx
			menuGroup := sys.Group("menu")
			// {
			// 	// menuGroup.GET("", handler.Menu().List)             // permission:menu:list
			// 	menuGroup.POST("", handler.Menu().Create)          // permission:menu:create
			// 	menuGroup.PUT("/:id", handler.Menu().Update)       // permission:menu:update
			// 	menuGroup.DELETE("/:ids", handler.Menu().Delete)   // permission:menu:delete
			// 	menuGroup.GET("/tree", handler.Menu().GetMenuTree) // permission:menu:tree
			// }
			{
				menuGroup.POST("", handler.SysMenu().Create)                   // system:menu:create
				menuGroup.PUT("/:id", handler.SysMenu().Update)                // system:menu:update
				menuGroup.DELETE("/:ids", handler.SysMenu().Delete)            // system:menu:delete
				menuGroup.GET("", handler.SysMenu().List)                      // system:menu:list
				menuGroup.GET("/tree", handler.SysMenu().GetMenuTree)          // system:menu:tree
				menuGroup.GET("/user-tree", handler.SysMenu().GetUserMenuTree) // system:menu:user-tree
			}

			// 字典管理
			dict := sys.Group("dict")
			{
				// 字典类型管理
				dictType := dict.Group("type")
				{
					dictType.POST("", handler.Dict().CreateDictType)        // system:dict:type:create
					dictType.PUT("/:id", handler.Dict().UpdateDictType)     // system:dict:type:update
					dictType.DELETE("/:ids", handler.Dict().DeleteDictType) // system:dict:type:delete
					dictType.GET("/:id", handler.Dict().GetDictType)        // system:dict:type:detail
					dictType.GET("", handler.Dict().ListDictType)           // system:dict:type:list
				}

				// 字典数据管理
				dictData := dict.Group("data")
				{
					dictData.POST("", handler.Dict().CreateDictData)              // system:dict:data:create
					dictData.PUT("/:id", handler.Dict().UpdateDictData)           // system:dict:data:update
					dictData.DELETE("/:ids", handler.Dict().DeleteDictData)       // system:dict:data:delete
					dictData.GET("/:id", handler.Dict().GetDictData)              // system:dict:data:detail
					dictData.GET("", handler.Dict().ListDictData)                 // system:dict:data:list
					dictData.GET("/type/:type", handler.Dict().GetDictDataByType) // system:dict:data:list:type
				}
			}

			// 系统菜单管理
			// menuGroup := sys.Group("menu")
			// {
			// 	menuGroup.POST("", handler.SysMenu().Create)        // system:menu:create
			// 	menuGroup.PUT("/:id", handler.SysMenu().Update)     // system:menu:update
			// 	menuGroup.DELETE("/:ids", handler.SysMenu().Delete) // system:menu:delete
			// 	menuGroup.GET("", handler.SysMenu().List)           // system:menu:list
			// }
		}
	}

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
