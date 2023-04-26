package coingecko

import (
	"fmt"
	"github.com/dl-nft-books/price-svc/internal/service/coingecko/models"
	"github.com/pkg/errors"
	"net/url"
	"time"
)

const (
	nftPriceEndpoint  = "/nfts/%s/contract/%s"
	nftPriceKeyFormat = "%s:%s:nft"
)

func (s *Service) GetNftPrice(platform, contract string) (string, error) {
	priceKey := fmt.Sprintf(nftPriceKeyFormat, platform, contract)

	if price, ok := s.priceCache[priceKey]; ok && price.ExpiresAt.Add(s.cacheExpiration).After(time.Now().UTC()) {
		fmt.Println("Get from cache")
		return price.Price, nil
	}

	parsedUrl, err := url.Parse(fmt.Sprintf(nftPriceEndpoint, platform, contract))
	if err != nil {
		return "", errors.Wrap(err, "failed to parse nft platforms url")

	}
	var response models.NftResponse
	err = s.connector.Get(parsedUrl, &response)
	if err != nil {
		return "", errors.Wrap(err, "failed to get nft platforms")
	}
	price := fmt.Sprintf("%f", response.FloorPrice.Usd)
	s.priceCache[priceKey] = Price{
		Price:     price,
		ExpiresAt: time.Now().UTC().Add(s.cacheExpiration),
	}
	fmt.Println(fmt.Sprintf("Set new price for %v in %v", platform, s.priceCache[priceKey].ExpiresAt.Format("02 Jan 2006 3:04:05 pm")))
	fmt.Println("CachePlatforms", s.priceCache)

	return price, nil
}
