package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/env"
	"github.com/sutd-statnlp/service-ladrawex/log"
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
	log.Debug("Init to load config files with environment:", env.EnvMode)
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
	log.Debug("Request to merge config files")
	err := loader.MergeInConfig()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

// Load loads all configurations from yaml files.
func Load() bool {
	log.Debug("Request to load config file")
	err := loader.ReadInConfig()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

// Watch watches changes from config files and reload.
func Watch() bool {
	log.Debug("Requet to watch changes from config file")
	loader.WatchConfig()
	loader.OnConfigChange(func(e fsnotify.Event) {
		log.Debug("Config file changed:", e.Name)
	})
	return loader != nil
}

// New reates new app configuration.
func New() *AppConfig {
	log.Debug("Request to create new app config")
	var config AppConfig
	err := loader.Unmarshal(&config)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &config
}

// Default returns the default app configuration.
func Default() *AppConfig {
	return appConfig
}
