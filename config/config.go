package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *viper.Viper
	mu   sync.RWMutex
	initOnce sync.Once
	initErr error
)

func Init(env string) {
	initOnce.Do(func() {
		initErr = initConfig(env)
	})
	if initErr != nil {
		log.Fatalf("Config initialization failed: %v", initErr)
	}
}

func initConfig(env string) error {
	v := viper.New()

	v.SetConfigName("default")
	v.SetConfigType("yaml")
	v.AddConfigPath("env/")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	envConfig := viper.New()
	envConfig.SetConfigName(env)
	envConfig.SetConfigType("yaml")
	envConfig.AddConfigPath("env/")

	if err := envConfig.ReadInConfig(); err != nil {
		return err
	}

	if err := v.MergeConfigMap(envConfig.AllSettings()); err != nil {
		return err
	}

	mu.Lock()
	cfg = v
	mu.Unlock()

	return nil
}

func GetConfig() *viper.Viper {
	mu.RLock()
	defer mu.RUnlock()

	if cfg == nil {
		log.Fatal("Config not initialized. Call config.Init() first.")
	}

	return cfg
}
