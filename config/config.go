package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Global *viper.Viper

func init() {
	configEnv := strings.ToLower(os.Getenv("CONFIG_ENV"))
	if configEnv == "" {
		configEnv = "development"
	}

	configFilePath := strings.ToLower(os.Getenv("CONFIG_FILE_PATH"))
	if configFilePath == "" {
		configFilePath = "./config"
	}

	configFileType := strings.ToLower(os.Getenv("CONFIG_FILE_TYPE"))
	if configFileType == "" {
		configFileType = "yaml"
	}

	Global = viper.New()

	Global.SetConfigName(configEnv)
	Global.AddConfigPath(configFilePath)
	Global.SetConfigType(configFileType)
	Global.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	Global.AutomaticEnv()

	if err := Global.ReadInConfig(); err != nil {
		log.Printf("Error unable to load config file: %v\n", err)
	}
}
