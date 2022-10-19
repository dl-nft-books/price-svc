package coingecko

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko/models"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/json-api-connector/cerrors"
)

const (
	priceEndpoint  = "/simple/token_price/%s"
	priceKeyFormat = "%s:%s:%s"
)

func (s *Service) GetPrice(platform, contract, vsAsset string) (string, error) {
	priceKey := fmt.Sprintf(priceKeyFormat, platform, contract, vsAsset)

	if price, ok := s.priceCache[priceKey]; ok && price.ExpiresAt.Add(s.cacheExpiration).After(time.Now().UTC()) {
		return price.Price, nil
	}

	var response models.PriceResponse

	parsedUrl, err := url.Parse(fmt.Sprintf(priceEndpoint, platform))
	if err != nil {
		return "", errors.Wrap(err, "failed to parse price url")
	}

	q := parsedUrl.Query()
	q.Set("contract_addresses", contract)
	q.Set("vs_currencies", vsAsset)

	parsedUrl.RawQuery = q.Encode()

	err = s.connector.Get(parsedUrl, &response)
	if err != nil {
		if cerr, ok := err.(cerrors.Error); ok && cerr.Status() == http.StatusNotFound {
			return "", nil
		}
		return "", errors.Wrap(err, "failed to get price")
	}

	price := response.GetPrice(contract, vsAsset)

	s.priceCache[priceKey] = Price{
		Price:     price,
		ExpiresAt: time.Now().UTC().Add(s.cacheExpiration),
	}

	return price, err
}
