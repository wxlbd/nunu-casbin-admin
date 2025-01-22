package main

import (
	"flag"
	"fmt"

	"github.com/wxlbd/gin-casbin-admin/cmd/server/wire"
	_ "github.com/wxlbd/gin-casbin-admin/docs" // 导入 swagger docs
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"go.uber.org/zap"
)

// @title Gin-Casbin-Admin API
// @version 1.0
// @description 基于 Gin + Casbin 的权限管理系统
// @termsOfService http://swagger.io/terms/

// @contact.name wxl
// @contact.url https://github.com/wxlbd
// @contact.email gopher095@gmail.com

// @license.name MIT
// @license.url https://github.com/wxlbd/gin-casbin-admin/blob/main/LICENSE

// @host localhost:8080
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	envConf := flag.String("conf", "configs/config.yaml", "config path, eg: -conf ./configs/config.yaml")
	flag.Parse()
	conf, err := config.NewConfig(*envConf)
	if err != nil {
		panic(err)
	}
	logger := log.NewLog(&conf.Log)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.Server.Host, conf.Server.Port)))
	if err = app.Run(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)); err != nil {
		panic(err)
	}
}
