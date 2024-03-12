package service

import (
	"context"

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
	generated.UnimplementedWeatherConditionServiceServer
}

func (p *Providers) GetWeather(ctx context.Context, request *generated.WeatherRequest) (*generated.WeatherResponse, error) {
	if ctx == nil {
		return nil, errors.New("weather request context is empty")
	}

	if request == nil {
		return nil, errors.New("weather request is empty")
	}

	hereMapsResponse, err := p.HereMapsClient.GetCoordinates(request.Location)
	if err != nil {
		return nil, err
	}

	openMateoResponse, err := p.OpenMateoClient.GetWeatherForecast(*hereMapsResponse)
	if err != nil {
		return nil, err
	}

	return openMateoResponse, nil
}

func NewWeatherDomainHandler(providers Providers) generated.WeatherConditionServiceServer {
	return &WeatherForecastServer{
		providers: providers,
	}
}
