package models

type CoordinatesResponse struct {
	Items []*Item `json:"items"`
}

type Item struct {
	Position *Position `json:"position"`
}

type Position struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}
