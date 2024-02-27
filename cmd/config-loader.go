package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initializeApplication() {
	properties = initializeConfiguration("./env/")
}

func initializeConfiguration(path string) *viper.Viper {
	viperConfigManager := viper.NewWithOptions(viper.KeyDelimiter("_"))
	viperConfigManager.SetConfigName("application")
	viperConfigManager.SetConfigType("yaml")
	viperConfigManager.AddConfigPath("/etc/config/")
	viperConfigManager.AddConfigPath(path)
	err := viperConfigManager.BindEnv()
	if err != nil {
		log.Warnf("Failed to bind a configuration key to the '%v , %v' environment variable with error %v", err)
	}
	viperConfigManager.AutomaticEnv()
	viperConfigManager.AllowEmptyEnv(true)

	viperConfigManager.WatchConfig()
	viperConfigManager.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file changed:", e.Name)
	})

	err = viperConfigManager.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("unable to start navigation-domain-handler due to missing application config %v", err))
	}
	log.Infof("loading application config from ", viperConfigManager.ConfigFileUsed())

	return viperConfigManager
}
