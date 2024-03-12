package openmateo

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"weather-conditions/service/providers/openmateo/models"

	heremaps "weather-conditions/service/providers/heremaps/models"
)

const (
	contentType        = "Content-Type"
	appType            = "application/json"
	envVarOpenMateoUrl = "OPEN_MATEO_URL"
	openMateoEndpoint  = "/v1/forecast"
	paramCurrent       = "current"
	paramCurrentValue  = "temperature_2m,is_day,rain,showers,snowfall"
	paramLatitude      = "latitude"
	paramLongitude     = "longitude"
)

type OpenMateoClient interface {
	GetWeatherForecast(heremaps.CoordinatesResponse) (*models.ForecastResponse, error)
}

type WeatherForecast struct {
	httClient *resty.Client
}

func (w *WeatherForecast) GetWeatherForecast(latLon heremaps.CoordinatesResponse) (*models.ForecastResponse, error) {
	url := w.httClient.BaseURL + openMateoEndpoint
	response, err := w.httClient.R().
		SetHeader(contentType, appType).
		SetQueryParam(paramCurrent, paramCurrentValue).
		SetQueryParam(paramLongitude, fmt.Sprintf("%f", latLon.Items[0].Position.Longitude)).
		SetQueryParam(paramLatitude, fmt.Sprintf("%f", latLon.Items[0].Position.Latitude)).
		SetResult(models.ForecastResponse{}).
		Get(url)
	if err != nil {
		log.Warnf("❌ OpenMateo API error: %s", err)
		return nil, err
	}

	return response.Result().(*models.ForecastResponse), nil
}

func NewOpenMateoClient(properties *viper.Viper) OpenMateoClient {
	url := properties.GetString(envVarOpenMateoUrl)
	client := resty.New()
	client.SetBaseURL(url)

	return &WeatherForecast{
		httClient: client,
	}
}
