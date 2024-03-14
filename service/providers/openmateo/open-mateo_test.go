package openmateo

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewOpenMateoClient(test *testing.T) {
	properties := viper.New()
	properties.Set(envVarOpenMateoUrl, mockUrl) // Use the same mock URL here

	emptyProperties := viper.New()

	tests := []struct {
		name string
		url  string
	}{
		{
			name: "all values set",
			url:  mockUrl,
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {
			tt.Parallel()
			client := NewOpenMateoClient(properties)
			assert.Equal(tt, t.url, client.(*WeatherForecast).httpClient.BaseURL)

			emptyClient := NewOpenMateoClient(emptyProperties)
			assert.Empty(tt, emptyClient.(*WeatherForecast).httpClient.BaseURL)
		})
	}
}
func TestGetWeatherForecast(test *testing.T) {

}

func TestGetWeatherForecastEmpty(test *testing.T) {

}

func TestGetWeatherForecastAll(test *testing.T) {

}
