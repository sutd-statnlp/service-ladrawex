package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/env"
	"github.com/sutd-statnlp/service-ladrawex/util/fileutil"
)

var (
	// loader is the configuration loader that loads config files.
	loader *viper.Viper

	// appConfig is app configuration that holds configured values.
	appConfig *AppConfig
)

// Init intialize viper config.
func init() {
	loader = NewLoader()
	Load(loader)

	if env.EnvMode == constant.ProdEnv {
		loader.SetConfigName(constant.ProdConfigName)
	} else {
		loader.SetConfigName(constant.DevConfigName)
	}
	Merge(loader)

	appConfig = New()
}

// Merge combines config files.
func Merge(inputLoader *viper.Viper) bool {
	err := inputLoader.MergeInConfig()
	if err != nil {
		return false
	}
	return true
}

// Load loads all configurations from yaml files.
func Load(inputLoader *viper.Viper) bool {
	err := inputLoader.ReadInConfig()
	if err != nil {
		return false
	}
	return true
}

// Watch watches changes from config files and reload.
func Watch(inputLoader *viper.Viper) bool {
	inputLoader.WatchConfig()
	inputLoader.OnConfigChange(func(e fsnotify.Event) {

	})
	return inputLoader != nil
}

// New reates new app configuration.
func New() *AppConfig {
	var config AppConfig
	err := loader.Unmarshal(&config)
	if err != nil {
		return nil
	}
	return &config
}

// Default returns the default app configuration.
func Default() *AppConfig {
	return appConfig
}

// NewLoader creates new loader.
func NewLoader() *viper.Viper {
	newLoader := viper.New()
	newLoader.SetConfigName(constant.ConfigName)
	newLoader.AddConfigPath(fileutil.FullPath(constant.ConfigPath))
	newLoader.SetConfigType(constant.ConfigType)
	return newLoader
}
