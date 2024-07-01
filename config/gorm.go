package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	username := Config.DBUsername
	password := Config.DBPassword
	host := Config.DBHost
	port := Config.DBPort
	database := Config.DBName

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	_, err = db.DB()
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	return db
}
