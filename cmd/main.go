package main

import (
	"github.com/spf13/viper"
	"weather-conditions/server"

	"weather-conditions/props"
)

var (
	properties *viper.Viper
	prop       *props.Properties
)

func main() {
	initializeApplication()

	startApplication()
}

func startApplication() {
	startServer()
}

func startServer() {
	providers := getProviders(properties)
	weatherConditionDomainHandler := NewWeatherConditionDomainHandler(providers, properties)
	servers := server.NewWeatherServer(properties, prop)
	//servers.ConfigureAPI(weatherConditionDomainHandler)
}

func getProviders(properties *viper.Viper) {

}
