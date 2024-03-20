package openmateo

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"weather-conditions/service/providers/heremaps/models"
)

const (
	mockPath             = "https://mockurl.com/v1/forecast"
	mockResponseOnlyTemp = "test-util/open-mateo-only-temperature.json"
	mockResponseAll      = "test-util/open-mateo-all.json"
)

var request = models.CoordinatesResponse{
	Items: []*models.Item{
		{
			Position: &models.Position{
				Latitude:  52.51604,
				Longitude: 13.37691,
			},
		},
	},
}

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
	restyClient := resty.New()
	jsonResponse, err := os.ReadFile(mockResponseOnlyTemp)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	recorder := httptest.NewRecorder()
	reader := bytes.NewReader(jsonResponse)
	_, err = recorder.Body.ReadFrom(reader)
	if err != nil {
		test.Fatal(err)
	}
	resp := recorder.Result()

	httpmock.RegisterResponder(http.MethodGet, mockPath, httpmock.ResponderFromResponse(resp))

	mockClient := OpenMateoMockClient(restyClient)
	result, err := mockClient.GetWeatherForecast(request)
	fmt.Println(result)

	// Assertions
	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.NotNil(test, result.TimeZone)
	assert.NotNil(test, result.TimeZoneUnit)
	assert.NotNil(test, result.CurrentUnits)
	assert.NotNil(test, result.CurrentValues)
	assert.NotNil(test, result.CurrentValues.Temperature)
	assert.Equal(test, float32(0.00), result.CurrentValues.Rain)
	assert.Equal(test, float32(0.00), result.CurrentValues.Showers)
	assert.Equal(test, float32(0.00), result.CurrentValues.Snowfall)
}

func TestGetWeatherForecastAll(test *testing.T) {
	restyClient := resty.New()
	jsonResponse, err := os.ReadFile(mockResponseAll)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	recorder := httptest.NewRecorder()
	reader := bytes.NewReader(jsonResponse)
	_, err = recorder.Body.ReadFrom(reader)
	if err != nil {
		test.Fatal(err)
	}
	resp := recorder.Result()

	httpmock.RegisterResponder(http.MethodGet, mockPath, httpmock.ResponderFromResponse(resp))

	mockClient := OpenMateoMockClient(restyClient)
	result, err := mockClient.GetWeatherForecast(request)
	fmt.Println(result)

	// Assertions
	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.NotNil(test, result.TimeZone)
	assert.NotNil(test, result.TimeZoneUnit)
	assert.NotNil(test, result.CurrentUnits)
	assert.NotNil(test, result.CurrentValues)
	assert.NotNil(test, result.CurrentValues.Temperature)
	assert.NotEqual(test, float32(0.00), result.CurrentValues.Rain)
	assert.NotEqual(test, float32(0.00), result.CurrentValues.Showers)
	assert.NotEqual(test, float32(0.00), result.CurrentValues.Snowfall)
}
