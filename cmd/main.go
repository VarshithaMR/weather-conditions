package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"weather-conditions/proto/generated"
	"weather-conditions/server"
	"weather-conditions/service"
	"weather-conditions/service/providers/heremaps"
	"weather-conditions/service/providers/openmateo"

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
	startgRPCServer()
}

func startgRPCServer() {
	// options like credentials, codec, keepalive params if required
	//opts := getServerOptions()
	providers := getProviders(properties)
	grpcServer := grpc.NewServer()
	servers := getWeatherServer()
	generated.RegisterWeatherConditionServiceServer(grpcServer,servers.weatherServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", prop.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

type weatherServices struct {
	weatherServer generated.WeatherConditionServiceServer
}

func getWeatherServer() weatherServices {

	return weatherServices{
		weatherServer:
	}
}

func getProviders(properties *viper.Viper) service.Providers {
	hereMaps := heremaps.NewHereMapsClient(properties)
	openMateo := openmateo.NewOpenMateoClient(properties)

	return service.Providers{
		HereMapsClient: hereMaps,
		OpenMateoClient: openMateo,
	}
}