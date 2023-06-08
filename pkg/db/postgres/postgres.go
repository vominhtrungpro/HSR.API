package postgres

import (
	"github.com/vominhtrungpro/config"
	"github.com/vominhtrungpro/pkg/db/basedb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Return new MySQL db instance
func New(c *config.MySQLConfig) (*gorm.DB, error) {
	db, err := basedb.Connect("mysql", *c)
	if err != nil {
		return nil, err
	}

	logMode := logger.Silent
	if c.Debug {
		logMode = logger.Error
	}

	return gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
}
