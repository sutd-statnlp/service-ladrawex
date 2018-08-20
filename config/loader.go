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
	loader = viper.New()
	loader.SetConfigName(constant.ConfigName)
	loader.AddConfigPath(fileutil.FullPath(constant.ConfigPath))
	loader.SetConfigType(constant.ConfigType)
	Load()

	if env.EnvMode == constant.ProdEnv {
		loader.SetConfigName(constant.ProdConfigName)
	} else {
		loader.SetConfigName(constant.DevConfigName)
	}
	Merge()

	appConfig = New()
}

// Merge combines config files.
func Merge() bool {
	err := loader.MergeInConfig()
	if err != nil {
		return false
	}
	return true
}

// Load loads all configurations from yaml files.
func Load() bool {
	err := loader.ReadInConfig()
	if err != nil {
		return false
	}
	return true
}

// Watch watches changes from config files and reload.
func Watch() bool {
	loader.WatchConfig()
	loader.OnConfigChange(func(e fsnotify.Event) {

	})
	return loader != nil
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
