package heremaps

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"weather-conditions/service/providers/heremaps/models"
)

const (
	contentType          = "Content-Type"
	appType              = "application/json"
	envVarHereMapsUrl    = "HERE_MAPS_URL"
	envVarHereMapsAPIKey = "HERE_MAPS_APIKEY"
	paramApiKey          = "apiKey"
	paramQuery           = "q"
	hereEndpoint         = "/v1/geocode"
)

type HereMapsClient interface {
	GetCoordinates(string) (*models.CoordinatesResponse, error)
}

type coordinatesRequest struct {
	httpClient *resty.Client
}

func (c *coordinatesRequest) GetCoordinates(query string) (*models.CoordinatesResponse, error) {
	url := c.httpClient.BaseURL + hereEndpoint
	response, err := c.httpClient.R().
		SetHeader(contentType, appType).
		SetQueryParam(paramQuery, query).
		SetResult(models.CoordinatesResponse{}).
		Get(url)

	if err != nil {
		log.Warnf("❌ HereMaps API error: %s", err)
		return nil, err
	}

	return response.Result().(*models.CoordinatesResponse), nil
}

func NewHereMapsClient(properties *viper.Viper) HereMapsClient {
	url := properties.GetString(envVarHereMapsUrl)
	apiKeyValue := properties.GetString(envVarHereMapsAPIKey)

	client := resty.New()
	client.SetBaseURL(url)
	client.SetQueryParam(paramApiKey, apiKeyValue)

	return &coordinatesRequest{
		httpClient: client,
	}
}
