package main

import (
	"flag"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/gormx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"gorm.io/gen"
)

func main() {
	envConf := flag.String("conf", "configs/config.yaml", "config path, eg: -conf ./configs/config.yaml")
	flag.Parse()
	conf, err := config.NewConfig(*envConf)
	if err != nil {
		panic(err)
	}
	logger := log.NewLog(&conf.Log)
	db := gormx.NewDB(conf, logger)
	// 启用 soft_delete 插件
	//db.Use(soft_delete.New(soft_delete.Config{
	//	Field: "deleted_at", // 软删除字段名
	//	Value: 1,            // 删除时的标志位值
	//}))
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/repository",
		ModelPkgPath: "internal/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})
	g.UseDB(db)

	// g.GenerateModel("dict_types")
	// g.GenerateModel("dict_data")
	g.ApplyBasic(g.GenerateModel("dict_types"), g.GenerateModel("dict_data"), model.Menu{}, model.Role{}, model.RoleMenus{}, model.User{}, model.UserRoles{})
	// g.GenerateAllTable()
	// g.GenerateAllTable()
	g.Execute()
}
