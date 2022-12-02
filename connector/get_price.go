package connector

import (
	"net/url"

	"github.com/pkg/errors"

	"gitlab.com/tokend/nft-books/price-svc/connector/models"
)

const (
	priceEndpoint = "pricer/price"
)

func (c *Connector) GetPrice(platform, contract string) (models.PriceResponse, error) {
	var response models.PriceResponse

	parsedUrl, err := url.Parse(priceEndpoint)
	if err != nil {
		return response, errors.Wrap(err, "failed to parse price url")
	}

	query := parsedUrl.Query()
	query.Set("platform", platform)
	query.Set("contract", contract)
	parsedUrl.RawQuery = query.Encode()

	err = c.Get(parsedUrl, &response)
	if err != nil {
		return response, errors.Wrap(err, "failed to get price")
	}

	return response, nil
}
