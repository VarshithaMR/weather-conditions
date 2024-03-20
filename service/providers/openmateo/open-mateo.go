package openmateo

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	heremaps "weather-conditions/service/providers/heremaps/models"
	"weather-conditions/service/providers/openmateo/models"
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
	httpClient *resty.Client
}

func (w *WeatherForecast) GetWeatherForecast(latLon heremaps.CoordinatesResponse) (*models.ForecastResponse, error) {
	baseUrl := w.httpClient.BaseURL
	response, err := w.httpClient.R().
		SetHeader(contentType, appType).
		SetQueryParam(paramCurrent, paramCurrentValue).
		SetQueryParam(paramLongitude, fmt.Sprintf("%f", latLon.Items[0].Position.Longitude)).
		SetQueryParam(paramLatitude, fmt.Sprintf("%f", latLon.Items[0].Position.Latitude)).
		SetResult(models.ForecastResponse{}).
		Get(baseUrl + openMateoEndpoint)
	if err != nil {
		log.Warnf("❌ OpenMateo API error: %s", err)
		return nil, err
	}

	// for mocking- test cases
	if baseUrl == mockUrl {
		var res models.ForecastResponse
		err = json.Unmarshal(response.Body(), &res)
		if err != nil {
			fmt.Println("Testcase : Unmarshal err", err)
		}
		return &res, nil
	}

	return response.Result().(*models.ForecastResponse), nil
}

func NewOpenMateoClient(properties *viper.Viper) OpenMateoClient {
	url := properties.GetString(envVarOpenMateoUrl)
	client := resty.New()
	client.SetBaseURL(url)

	return &WeatherForecast{
		httpClient: client,
	}
}
