package main

import (
	"flag"
	"k-style-test/app"
	"k-style-test/config"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	var logFile *os.File
	LoadConfig("config/config.yml", config.Config.Log.FileNamePrefix, config.Config.Log.Level, logFile)

	db := config.NewDatabase()
	validator := validator.New()

	echo := echo.New()
	// gin := gin.Default()
	// echo.Use(middleware.Logger())

	app.InitApp(&app.Application{
		DB:            db,
		Validate:      validator,
		Echo:          echo,
		WsConnections: make([]*config.WebSocketConnection, 0),
	})

	echo.Start(":" + config.Config.Port)
	logrus.Infof("App Running... Url: %v Port: %v", config.Config.AppUrl, config.Config.Port)

}

// LoadConfig to load config from config.yml file
func LoadConfig(configPath, logFileName, logLevel string, logFile *os.File) {
	configFile := flag.String("config", configPath, "main configuration file")
	config.InitConfig(configFile)
	flag.Parse()
	logrus.Infof("Reads configuration from %s", *configFile)
	config.InitLog(config.Config.Log.FileNamePrefix, logLevel, logFile)
}
