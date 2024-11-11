package dbconnect

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/pkg/config"
)

type DBConnect struct {
	config *config.Config
}

func NewDBConnnect(config *config.Config) DBConnect {
	return DBConnect{
		config: config,
	}
}

func (dbConnect DBConnect) GetDBConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConnect.config.DbConfig.DbHost,
		dbConnect.config.DbConfig.DbUser,
		dbConnect.config.DbConfig.DbPassword,
		dbConnect.config.DbConfig.DbName,
		dbConnect.config.DbConfig.DbPort,
		dbConnect.config.DbConfig.DbSllMode,
		dbConnect.config.DbConfig.DbTimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
