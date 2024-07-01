package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	yaml "github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
)

var Config Configuration

type Configuration struct {
	Env        string `yaml:"env"`
	Port       string `yaml:"port"`
	DBList     string `yaml:"dbList"`
	DBUsername string `yaml:"dbUsername"`
	DBPassword string `yaml:"dbPassword"`
	DBPort     string `yaml:"dbPort"`
	DBHost     string `yaml:"dbHost"`
	DBName     string `yaml:"dbName"`
	AppName    string `yaml:"appName"`
	AppUrl     string `yaml:"appUrl"`
	JwtSecret  string `yaml:"jwtSecret"`
	Log        struct {
		FileNamePrefix string `yaml:"filenamePrefix"`
		Level          string `yaml:"level"`
	} `yaml:"log"`
}

func InitConfig(fn *string) {

	if err := LoadYAML(fn, &Config); err != nil {
		log.Errorf("LoadConfigFromFile() - Failed opening config file %s\n%s", &fn, err)
		os.Exit(1)
	}
	//log.Logf("Loaded configs: %v", Param)
	t := time.Now()
	sDate := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
	if Config.Env == "local" {
		Config.Log.FileNamePrefix = Config.Log.FileNamePrefix + sDate + ".log"
	} else {
		Config.Log.FileNamePrefix = Config.Log.FileNamePrefix + ".log"
	}
}

func LoadYAML(filename *string, v interface{}) error {
	raw, err := ioutil.ReadFile(*filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, v)
	if err != nil {
		return err
	}
	return nil
}
