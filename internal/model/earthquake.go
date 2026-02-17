package model

import "time"

type Earthquake struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	Magnitude float64   `json:"magnitude"`
	Depth     float64   `json:"depth"`
	Time      time.Time `json:"time"`
	Link      string    `json:"link"`
}

// RSS Structures for TMD Earthquake
type EarthquakeRSS struct {
	Channel struct {
		Items []EarthquakeItem `xml:"item"`
	} `xml:"channel"`
}

type EarthquakeItem struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description string  `xml:"description"`
	PubDate     string  `xml:"pubDate"`
	Lat         float64 `xml:"http://www.w3.org/2003/01/geo/ lat"`
	Long        float64 `xml:"http://www.w3.org/2003/01/geo/ long"`
	Magnitude   float64 `xml:"http://www.earthquake.tmd.go.th magnitude"`
	Depth       float64 `xml:"http://www.earthquake.tmd.go.th depth"`
}
