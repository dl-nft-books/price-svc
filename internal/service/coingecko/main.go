package coingecko

import (
	"time"

	"github.com/dl-nft-books/price-svc/internal/service/helpers"
)

type Service struct {
	connector       *helpers.Connector
	priceCache      PriceCache
	cacheExpiration time.Duration
}

func NewService(connector *helpers.Connector, cacheExpiration time.Duration) *Service {
	return &Service{
		connector:       connector,
		priceCache:      make(PriceCache),
		cacheExpiration: cacheExpiration,
	}
}
