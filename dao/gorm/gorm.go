package gorm

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hinccvi/go-template/config"
	"hinccvi/go-template/log"
)

var DB *gorm.DB

func Init(env string) {
	user := config.Conf.DBConfig.User
	password := config.Conf.DBConfig.Password
	dbname := config.Conf.DBConfig.DBName
	host := config.Conf.DBConfig.Host
	port := config.Conf.DBConfig.Port
	var err error
	var logger log.GormLogger

	if env == "dev" {
		logger = log.New()
		logger.SetAsDefault()
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", host, user, password, dbname, port),
	}), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		log.Panic("Fail to connect DB", zap.Error(err))
	}

	log.Info("Gorm successfully init")
}
