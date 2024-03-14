package heremaps

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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

}

func TestGetCoordinatesEmpty(test *testing.T) {

}
