package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	config = viper.New()

	config.SetConfigName("default")
	config.SetConfigType("yaml")
	config.AddConfigPath("env/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading default config file: %v", err)
	}

	envConfig := viper.New()
	envConfig.SetConfigName(env)
	envConfig.SetConfigType("yaml")
	envConfig.AddConfigPath("env/")

	if err := envConfig.ReadInConfig(); err != nil {
		log.Fatalf("Error reading %s config file: %v", env, err)
	}

	if err := config.MergeConfigMap(envConfig.AllSettings()); err != nil {
		log.Fatalf("Error merging %s config: %v", env, err)
	}
}

func GetConfig() *viper.Viper {
	if config == nil {
		log.Fatal("Config not initialized")
	}

	return config
}
