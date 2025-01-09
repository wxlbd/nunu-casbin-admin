package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/nunu-casbin-admin/internal/handler"
	"github.com/wxlbd/nunu-casbin-admin/internal/middleware"
	"github.com/wxlbd/nunu-casbin-admin/pkg/config"
	"github.com/wxlbd/nunu-casbin-admin/pkg/log"
)

func NewServerHTTP(
	cfg *config.Config,
	logger *log.Logger,
	userHandler *handler.Handler,
) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.POST("/login", userHandler.User().Login)
	r.POST("/logout", userHandler.User().Logout)
	r.POST("/refresh", userHandler.User().RefreshToken)

	return r
}
