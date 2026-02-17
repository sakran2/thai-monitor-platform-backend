package collector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type TMDClient struct {
	Token      string
	BaseURL    string
	HTTPClient *http.Client
}

type TMDAPIResponse struct {
	WeatherForecasts []TMDLocationWrapper `json:"WeatherForecasts"`
}

type TMDLocationWrapper struct {
	Location  TMDLocInfo    `json:"location"`
	Forecasts []TMDForecast `json:"forecasts"`
}

type TMDLocInfo struct {
	Province string  `json:"province"`
	AreaType string  `json:"areatype"`
	Amphoe   string  `json:"amphoe"`
	Region   string  `json:"region"`
	Geocode  string  `json:"geocode"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

type TMDForecast struct {
	Time string                 `json:"time"`
	Data map[string]interface{} `json:"data"`
}

func NewTMDClient(token string) *TMDClient {
	return &TMDClient{
		Token:   token,
		BaseURL: "https://data.tmd.go.th/nwpapi/v1/forecast/location/hourly/place",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *TMDClient) FetchForecast(province, amphoe string, date string, hour int, duration int) (*TMDAPIResponse, error) {
	params := url.Values{}
	params.Set("province", province)
	if amphoe != "" {
		params.Set("amphoe", amphoe)
	}
	if date != "" {
		params.Set("date", date)
	}
	if hour >= 0 {
		params.Set("hour", fmt.Sprintf("%d", hour))
	}
	params.Set("fields", "tc,rh,rain,ws10m,wd10m,cond")
	params.Set("duration", fmt.Sprintf("%d", duration))

	reqURL := fmt.Sprintf("%s?%s", c.BaseURL, params.Encode())

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch forecast: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("TMD API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result TMDAPIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w (body: %s)", err, string(body))
	}

	return &result, nil
}
