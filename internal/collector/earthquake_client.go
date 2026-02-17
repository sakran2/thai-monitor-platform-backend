package collector

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"thai-monitor-platform/backend/internal/model"
)

const EarthquakeRSSURL = "https://earthquake.tmd.go.th/feed/rss_tmd.xml"

func FetchLatestEarthquakes() ([]model.Earthquake, error) {
	resp, err := http.Get(EarthquakeRSSURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch earthquake rss: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("earthquake rss returned status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read earthquake rss body: %v", err)
	}

	var rss model.EarthquakeRSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal earthquake rss: %v", err)
	}

	earthquakes := make([]model.Earthquake, 0)
	for _, item := range rss.Channel.Items {
		eq := model.Earthquake{
			ID:        item.Link, // Use link as a simple ID
			Title:     item.Title,
			Lat:       item.Lat,
			Long:      item.Long,
			Magnitude: item.Magnitude,
			Depth:     item.Depth,
			Link:      item.Link,
		}
		// In a real app, we would parse item.PubDate into time.Time
		earthquakes = append(earthquakes, eq)
	}

	return earthquakes, nil
}
