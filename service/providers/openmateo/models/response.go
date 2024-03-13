package models

type ForecastResponse struct {
	TimeZoneUnit  string         `json:"timezone_abbreviation"`
	TimeZone      string         `json:"timezone"`
	CurrentUnits  *CurrentUnits  `json:"current_units,omitempty"`
	CurrentValues *CurrentValues `json:"current,omitempty"`
}

type CurrentUnits struct {
	Temperature string `json:"temperature_2m"`
	Rain        string `json:"rain,omitempty"`
	Showers     string `json:"showers,omitempty"`
	Snowfall    string `json:"snowfall,omitempty"`
}

type CurrentValues struct {
	Temperature float32 `json:"temperature_2m"`
	Rain        float32 `json:"rain,omitempty"`
	Showers     float32 `json:"showers,omitempty"`
	Snowfall    float32 `json:"snowfall,omitempty"`
}
