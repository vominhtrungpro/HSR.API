package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/vominhtrungpro/config"
	"github.com/vominhtrungpro/pkg/db/basedb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(c *config.MySQLConfig) (*gorm.DB, error) {
	db, err := basedb.Connect("mysql", *c)
	if err != nil {
		return nil, err
	}

	logMode := logger.Silent
	if c.Debug {
		logMode = logger.Error
	}

	return gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
}
