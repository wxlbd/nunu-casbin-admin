package main

import (
	"flag"
	"fmt"

	"github.com/wxlbd/gin-casbin-admin/cmd/server/wire"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"go.uber.org/zap"
)

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
