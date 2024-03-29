package connector

import (
	"github.com/pkg/errors"
	"net/url"

	"github.com/dl-nft-books/price-svc/connector/models"
)

const (
	nftPriceEndpoint = "pricer/nft"
)

func (c *Connector) GetNftPrice(platform, contract string) (models.NftPriceResponse, error) {
	var response models.NftPriceResponse

	parsedUrl, err := url.Parse(nftPriceEndpoint)
	if err != nil {
		return response, errors.Wrap(err, "failed to parse nft price url")
	}

	query := parsedUrl.Query()
	query.Set("platform", platform)
	query.Set("contract", contract)
	parsedUrl.RawQuery = query.Encode()

	err = c.Get(parsedUrl, &response)
	if err != nil {
		return response, errors.Wrap(err, "failed to get nft price")
	}

	return response, nil
}
