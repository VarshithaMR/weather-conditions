package service

import (
	"context"
)

func (s *ser.Serve) GetWeather(context.Context, *WeatherRequest) (*WeatherResponse, error)
