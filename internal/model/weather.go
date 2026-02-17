package model

import "time"

// WeatherForecast represents a weather forecast record
type WeatherForecast struct {
	ID           int       `json:"id"`
	Province     string    `json:"province"`
	Amphoe       string    `json:"amphoe"`
	Geocode      string    `json:"geocode"`
	Region       string    `json:"region"`
	Lat          float64   `json:"lat"`
	Lon          float64   `json:"lon"`
	ForecastTime time.Time `json:"forecast_time"`
	Tc           *float64  `json:"tc"`
	TcMax        *float64  `json:"tc_max"`
	TcMin        *float64  `json:"tc_min"`
	Rh           *float64  `json:"rh"`
	Rain         *float64  `json:"rain"`
	Ws10m        *float64  `json:"ws10m"`
	Wd10m        *float64  `json:"wd10m"`
	Slp          *float64  `json:"slp"`
	Cond         *int      `json:"cond"`
	CreatedAt    time.Time `json:"created_at"`
}

// TMD API Response structures
type TMDResponse struct {
	WeatherForecast TMDWeatherForecast `json:"WeatherForecasts"`
}

type TMDWeatherForecast struct {
	Locations []TMDLocation `json:"location"`
}

type TMDLocation struct {
	Location  TMDLocationInfo `json:"location"`
	Forecasts []TMDForecast   `json:"forecasts"`
}

type TMDLocationInfo struct {
	Province string  `json:"province"`
	AreaType string  `json:"areatype"`
	Amphoe   string  `json:"amphoe"`
	Region   string  `json:"region"`
	Geocode  string  `json:"geocode"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

type TMDForecast struct {
	Time string      `json:"time"`
	Data TMDForecastData `json:"data"`
}

type TMDForecastData struct {
	Tc    *float64 `json:"tc"`
	TcMax *float64 `json:"tc_max"`
	TcMin *float64 `json:"tc_min"`
	Rh    *float64 `json:"rh"`
	Rain  *float64 `json:"rain"`
	Ws10m *float64 `json:"ws10m"`
	Wd10m *float64 `json:"wd10m"`
	Slp   *float64 `json:"slp"`
	Cond  *int     `json:"cond"`
}

// WeatherCondition maps condition code to description
var WeatherConditions = map[int]string{
	1:  "ท้องฟ้าแจ่มใส",
	2:  "มีเมฆบางส่วน",
	3:  "เมฆเป็นส่วนมาก",
	4:  "มีเมฆมาก",
	5:  "ฝนตกเล็กน้อย",
	6:  "ฝนปานกลาง",
	7:  "ฝนตกหนัก",
	8:  "ฝนฟ้าคะนอง",
	9:  "อากาศหนาวจัด",
	10: "อากาศหนาว",
	11: "อากาศเย็น",
	12: "อากาศร้อนจัด",
}

var WeatherConditionsEN = map[int]string{
	1:  "Clear",
	2:  "Partly cloudy",
	3:  "Cloudy",
	4:  "Overcast",
	5:  "Light rain",
	6:  "Moderate rain",
	7:  "Heavy rain",
	8:  "Thunderstorm",
	9:  "Very cold",
	10: "Cold",
	11: "Cool",
	12: "Very hot",
}

// ForecastResponse is the API response format
type ForecastResponse struct {
	Province     string          `json:"province"`
	Amphoe       string          `json:"amphoe"`
	Lat          float64         `json:"lat"`
	Lon          float64         `json:"lon"`
	Forecasts    []ForecastItem  `json:"forecasts"`
}

type ForecastItem struct {
	ForecastTime string   `json:"forecast_time"`
	Tc           *float64 `json:"tc"`
	TcMax        *float64 `json:"tc_max"`
	TcMin        *float64 `json:"tc_min"`
	Rh           *float64 `json:"rh"`
	Rain         *float64 `json:"rain"`
	Ws10m        *float64 `json:"ws10m"`
	Wd10m        *float64 `json:"wd10m"`
	Cond         *int     `json:"cond"`
	CondText     string   `json:"cond_text"`
	CondTextEN   string   `json:"cond_text_en"`
}
