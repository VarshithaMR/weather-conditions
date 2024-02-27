package main

import (
	"github.com/spf13/viper"

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
	weatherConditionDomainHandler := NewWestherConditionDomainHandler(providers, properties)
	server := NewWeatherServer(prop, properties)
	server.ConfigureApi(weatherConditionDomainHandler)
}
