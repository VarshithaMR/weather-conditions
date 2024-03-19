package heremaps

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
)

const (
	mockPath                  = "https://mockurl.com/v1/geocode"
	mockResponse              = "test-util/here-maps-coordinates.json"
	mockResponseNoCoordinates = "test-util/here-maps-no-coordinates.json"
)

func TestNewHereMapsClient(test *testing.T) {
	properties := viper.New()
	properties.Set(envVarHereMapsUrl, mockUrl) // Use the same mock URL here
	properties.Set(envVarHereMapsAPIKey, mockApiKey)

	emptyProperties := viper.New()

	tests := []struct {
		name string
		url  string
		key  string
	}{
		{
			name: "all values set",
			url:  mockUrl,
			key:  mockApiKey,
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(tt *testing.T) {
			tt.Parallel()
			client := NewHereMapsClient(properties)
			assert.Equal(tt, t.url, client.(*coordinatesRequest).httpClient.BaseURL)
			assert.Equal(tt, t.key, client.(*coordinatesRequest).httpClient.QueryParam.Get(paramApiKey))

			emptyClient := NewHereMapsClient(emptyProperties)
			assert.Empty(tt, emptyClient.(*coordinatesRequest).httpClient.BaseURL)
			assert.Empty(tt, emptyClient.(*coordinatesRequest).httpClient.QueryParam.Get(paramApiKey))
		})
	}
}
func TestGetCoordinates(test *testing.T) {
	restyClient := resty.New()
	jsonResponse, err := os.ReadFile(mockResponse)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	recorder := httptest.NewRecorder()
	reader := bytes.NewReader(jsonResponse)
	x, err := recorder.Body.ReadFrom(reader)
	fmt.Println(x)
	if err != nil {
		test.Fatal(err)
	}
	resp := recorder.Result()

	httpmock.RegisterResponder(http.MethodGet, mockPath, httpmock.ResponderFromResponse(resp))

	mockHereMapsClient := HereMapsMockClient(restyClient)
	result, err := mockHereMapsClient.GetCoordinates("Berlin")
	fmt.Println(result)

	// Assertions
	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.NotNil(test, result.Items)
	assert.NotNil(test, result.Items[0].Position)
	assert.NotNil(test, result.Items[0].Position.Longitude)
	assert.NotNil(test, result.Items[0].Position.Latitude)
}

func TestGetCoordinatesEmpty(test *testing.T) {
	restyClient := resty.New()
	jsonResponse, err := os.ReadFile(mockResponseNoCoordinates)
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

	mockHereMapsClient := HereMapsMockClient(restyClient)
	result, err := mockHereMapsClient.GetCoordinates("Berlin")
	fmt.Println(result)

	// Assertions
	assert.NoError(test, err)
	assert.NotNil(test, result)
	assert.Nil(test, result.Items)
}
