package main

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"weather-conditions/props"
	"weather-conditions/proto/generated"
	"weather-conditions/service"
	"weather-conditions/service/providers/heremaps"
	"weather-conditions/service/providers/openmateo"
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

	// create new gRPC server
	grpcServer := grpc.NewServer()

	//initialise Weather server
	servers := getWeatherServer()

	// register your server
	generated.RegisterWeatherConditionServiceServer(grpcServer, servers.weatherServer)

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
	//get all external providers
	providers := getProviders(properties)

	return weatherServices{
		weatherServer: service.NewWeatherDomainHandler(providers),
	}
}

func getProviders(properties *viper.Viper) service.Providers {
	hereMaps := heremaps.NewHereMapsClient(properties)
	openMateo := openmateo.NewOpenMateoClient(properties)

	return service.Providers{
		HereMapsClient:  hereMaps,
		OpenMateoClient: openMateo,
	}
}
