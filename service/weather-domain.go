package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"weather-conditions/proto/generated"
	"weather-conditions/service/providers/heremaps"
	"weather-conditions/service/providers/openmateo"
)

type Providers struct {
	HereMapsClient  heremaps.HereMapsClient
	OpenMateoClient openmateo.OpenMateoClient
}

type WeatherForecastServer struct {
	providers Providers
	Host      string
	Path      string
	Port      int
	generated.UnimplementedWeatherConditionServiceServer
}

func (w *WeatherForecastServer) GetWeather(ctx context.Context, request *generated.WeatherRequest) (*generated.WeatherResponse, error) {
	if ctx == nil {
		return nil, errors.New("weather request context is empty")
	}

	if request == nil {
		return nil, errors.New("weather request is empty")
	}

	hereMapsResponse, err := w.providers.HereMapsClient.GetCoordinates(request.Location)
	if err != nil {
		return nil, err
	}

	openMateoResponse, err := w.providers.OpenMateoClient.GetWeatherForecast(*hereMapsResponse)
	if err != nil {
		return nil, err
	}

	temperature := fmt.Sprintf("%f", openMateoResponse.CurrentValues.Temperature) + " " + openMateoResponse.CurrentUnits.Temperature
	timezone := openMateoResponse.TimeZone + " " + openMateoResponse.TimeZoneUnit
	x := &generated.WeatherResponse{
		Temperature: temperature,
		Timezone:    timezone,
	}

	return x, nil
}

type ServerOption func(*WeatherForecastServer)

func NewWeatherDomainHandler(providers Providers, options ...ServerOption) generated.WeatherConditionServiceServer {
	weatherServer := &WeatherForecastServer{
		providers: providers,
	}

	for _, option := range options {
		if option != nil {
			option(weatherServer)
		}
	}
	return weatherServer
}

func WithHost(host string) ServerOption {
	return func(w *WeatherForecastServer) {
		w.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(w *WeatherForecastServer) {
		w.Port = port
	}
}

func WithPath(path string) ServerOption {
	return func(w *WeatherForecastServer) {
		w.Path = path
	}
}
