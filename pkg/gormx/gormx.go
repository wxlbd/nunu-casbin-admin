package gormx

import (
	"time"

	"github.com/glebarez/sqlite"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"github.com/wxlbd/gin-casbin-admin/pkg/zapgorm2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(conf *config.Config, l *log.Logger) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	logger := zapgorm2.New(l.Logger)
	driver := conf.Database.Driver
	dsn := conf.Database.DSN

	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html
	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger,
		})
	case "postgres":
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		panic("unknown db driver")
	}
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
