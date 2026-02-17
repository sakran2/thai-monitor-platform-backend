package database

import (
	"database/sql"
	"fmt"
	"log"

	"thai-monitor-platform/backend/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(cfg *config.Config) error {
	dsn := cfg.GetDSN()
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return nil
}

func Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS weather_forecasts (
		id SERIAL PRIMARY KEY,
		province VARCHAR(100) NOT NULL,
		amphoe VARCHAR(100),
		geocode VARCHAR(20),
		region VARCHAR(10),
		lat DOUBLE PRECISION,
		lon DOUBLE PRECISION,
		forecast_time TIMESTAMPTZ NOT NULL,
		tc_max DOUBLE PRECISION,
		tc_min DOUBLE PRECISION,
		rh DOUBLE PRECISION,
		rain DOUBLE PRECISION,
		ws10m DOUBLE PRECISION,
		wd10m DOUBLE PRECISION,
		slp DOUBLE PRECISION,
		cond INTEGER,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		UNIQUE(province, amphoe, forecast_time)
	);

	CREATE INDEX IF NOT EXISTS idx_weather_province ON weather_forecasts(province);
	CREATE INDEX IF NOT EXISTS idx_weather_forecast_time ON weather_forecasts(forecast_time);
	`

	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	log.Println("✅ Database migrated")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
