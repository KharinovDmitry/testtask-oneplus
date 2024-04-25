package dto

import "time"

type CourseDTO struct {
	ID                       string    `json:"id"`
	Symbol                   string    `json:"symbol"`
	Name                     string    `json:"name"`
	Image                    string    `json:"image"`
	CurrentPrice             float64   `json:"current_price"`
	MarketCap                int       `json:"market_cap"`
	MarketCapRank            int       `json:"market_cap_rank"`
	FullyDilutedValuation    int       `json:"fully_diluted_valuation"`
	TotalVolume              float64   `json:"total_volume"`
	High24h                  float64   `json:"high_24h"`
	Low24h                   float64   `json:"low_24h"`
	PriceChange24h           float64   `json:"price_change_24h"`
	PriceChangePercentage24h float64   `json:"price_change_percentage_24h"`
	PercentChange24h         float64   `json:"percent_change_24h"`
	CirculatingSupply        float64   `json:"circulating_supply"`
	TotalSupply              float64   `json:"total_supply"`
	MaxSupply                float64   `json:"max_supply"`
	Ath                      float64   `json:"ath"`
	AthChangePercentage      float64   `json:"ath_change_percentage"`
	AthDate                  time.Time `json:"ath_date"`
	Atl                      float64   `json:"atl"`
	AtlChangePercentage      float64   `json:"atl_change_percentage"`
	AtlDate                  time.Time `json:"atl_date"`
	Roi                      RoiDTO    `json:"roi"`
	LastUpdated              time.Time `json:"last_updated"`
}

type RoiDTO struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}
