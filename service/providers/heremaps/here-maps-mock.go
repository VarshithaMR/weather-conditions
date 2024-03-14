package heremaps

import "github.com/go-resty/resty/v2"

const (
	mockUrl    = "https://mockurl.com"
	mockApiKey = "mockapikey"
)

func HereMapsMockClient(mapsClient *resty.Client) HereMapsClient {
	mapsClient.SetBaseURL(mockUrl)
	mapsClient.SetQueryParam(paramApiKey, mockApiKey)

	return &coordinatesRequest{
		httpClient: mapsClient,
	}
}
