package service

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"thai-monitor-platform/backend/internal/collector"
	"thai-monitor-platform/backend/internal/config"
	"thai-monitor-platform/backend/internal/model"
)

type WeatherService struct {
	tmdClient *collector.TMDClient
}

func NewWeatherService(cfg *config.Config) *WeatherService {
	return &WeatherService{
		tmdClient: collector.NewTMDClient(cfg.TMDToken),
	}
}

func (s *WeatherService) GetForecast(province, amphoe string, duration int) (*model.ForecastResponse, error) {
	if duration <= 0 {
		duration = 24 // 24 hours for hourly
	}
	if duration > 120 {
		duration = 120 // Max 5 days of hourly
	}

	now := time.Now()
	dateStr := now.Format("2006-01-02")
	hour := now.Hour()

	resp, err := s.tmdClient.FetchForecast(province, amphoe, dateStr, hour, duration)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from TMD: %w", err)
	}

	if len(resp.WeatherForecasts) == 0 {
		return nil, fmt.Errorf("no forecast data found for province: %s", province)
	}

	loc := resp.WeatherForecasts[0]
	result := &model.ForecastResponse{
		Province: loc.Location.Province,
		Amphoe:   loc.Location.Amphoe,
		Lat:      loc.Location.Lat,
		Lon:      loc.Location.Lon,
	}

	for _, fc := range loc.Forecasts {
		item := model.ForecastItem{
			ForecastTime: fc.Time,
		}

		// Parse data fields from map[string]interface{}
		if v, ok := fc.Data["tc"]; ok {
			f := toFloat64(v)
			item.Tc = &f
		}
		if v, ok := fc.Data["tc_max"]; ok {
			f := toFloat64(v)
			item.TcMax = &f
		}
		if v, ok := fc.Data["tc_min"]; ok {
			f := toFloat64(v)
			item.TcMin = &f
		}
		if v, ok := fc.Data["rh"]; ok {
			f := toFloat64(v)
			item.Rh = &f
		}
		if v, ok := fc.Data["rain"]; ok {
			f := toFloat64(v)
			item.Rain = &f
		}
		if v, ok := fc.Data["ws10m"]; ok {
			f := toFloat64(v)
			item.Ws10m = &f
		}
		if v, ok := fc.Data["wd10m"]; ok {
			f := toFloat64(v)
			item.Wd10m = &f
		}
		if v, ok := fc.Data["cond"]; ok {
			c := toInt(v)
			item.Cond = &c
			if text, exists := model.WeatherConditions[c]; exists {
				item.CondText = text
			}
			if text, exists := model.WeatherConditionsEN[c]; exists {
				item.CondTextEN = text
			}
		}

		result.Forecasts = append(result.Forecasts, item)
	}

	log.Printf("📡 Fetched hourly forecast for %s %s: %d slots", province, amphoe, len(result.Forecasts))
	return result, nil
}

func toFloat64(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	default:
		return 0
	}
}

func toInt(v interface{}) int {
	switch val := v.(type) {
	case float64:
		return int(val)
	case int:
		return val
	case string:
		i, _ := strconv.Atoi(val)
		return i
	default:
		return 0
	}
}
