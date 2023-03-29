package connector

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/dl-nft-books/price-svc/connector/models"
)

const (
	priceEndpoint = "pricer/price"
)

func (c *Connector) GetPrice(platform, contract string, chainId int64) (models.PriceResponse, error) {
	var response models.PriceResponse

	parsedUrl, err := url.Parse(priceEndpoint)
	if err != nil {
		return response, errors.Wrap(err, "failed to parse price url")
	}

	query := parsedUrl.Query()
	query.Set("platform", platform)
	query.Set("contract", contract)
	query.Set("chain_id", strconv.FormatInt(chainId, 10))
	parsedUrl.RawQuery = query.Encode()

	err = c.Get(parsedUrl, &response)
	if err != nil {
		return response, errors.Wrap(err, "failed to get price")
	}

	return response, nil
}
