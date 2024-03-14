package openmateo

import "github.com/go-resty/resty/v2"

const (
	mockUrl = "https://mockurl.com"
)

func OpenMateoMockClient(forecastClient *resty.Client) OpenMateoClient {
	forecastClient.SetBaseURL(mockUrl)

	return &WeatherForecast{
		httClient: forecastClient,
	}
}
