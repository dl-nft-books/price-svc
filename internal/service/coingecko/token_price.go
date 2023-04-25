package coingecko

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dl-nft-books/price-svc/internal/service/coingecko/models"
	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/json-api-connector/cerrors"
)

const (
	tokenPriceEndpoint = "/simple/token_price/%s"
	coinPriceEndpoint  = "/simple/price"
	priceKeyFormat     = "%s:%s:%s"
	polygonId          = "polygon-pos"
	nativeMaticId      = "matic-network"
)

func (s *Service) GetTokenContractInfo(address common.Address) {
}

func (s *Service) GetPrice(platform, contract, vsAsset string) (string, error) {
	priceKey := fmt.Sprintf(priceKeyFormat, platform, contract, vsAsset)

	if price, ok := s.priceCache[priceKey]; ok && price.ExpiresAt.Add(s.cacheExpiration).After(time.Now().UTC()) {
		return price.Price, nil
	}

	var (
		price string
		err   error
	)

	if contract != "" {
		price, err = s.getPriceContract(platform, contract, vsAsset)
	} else {
		price, err = s.getPriceNative(platform, vsAsset)
	}

	if err != nil {
		return "", err
	}

	s.priceCache[priceKey] = Price{
		Price:     price,
		ExpiresAt: time.Now().UTC().Add(s.cacheExpiration),
	}

	return price, nil
}

func (s *Service) getPriceContract(platform, contract, vsAsset string) (string, error) {
	var response models.PriceResponse

	parsedUrl, err := url.Parse(fmt.Sprintf(tokenPriceEndpoint, platform))
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

	return price, err
}

func (s *Service) getPriceNative(platform, vsAsset string) (string, error) {
	var response models.PriceResponse

	parsedUrl, err := url.Parse(coinPriceEndpoint)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse price url")
	}

	// Changing platform ids to get MATIC price in $
	if platform == polygonId {
		platform = nativeMaticId
	}

	q := parsedUrl.Query()
	q.Set("ids", platform)
	q.Set("vs_currencies", vsAsset)

	parsedUrl.RawQuery = q.Encode()
	err = s.connector.Get(parsedUrl, &response)
	if err != nil {
		if cerr, ok := err.(cerrors.Error); ok && cerr.Status() == http.StatusNotFound {
			return "", nil
		}
		return "", errors.Wrap(err, "failed to get price")
	}

	price := response.GetPrice(platform, vsAsset)

	return price, err
}
