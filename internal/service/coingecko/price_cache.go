package coingecko

import "time"

type Price struct {
	Price     string
	ExpiresAt time.Time
}

type PriceCache map[string]Price
