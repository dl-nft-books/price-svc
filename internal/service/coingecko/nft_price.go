package coingecko

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko/models"
	"net/url"
)

const (
	nftPriceEndpoint = "/nfts/%s/contract/%s"
)

func (s *Service) GetNftPrice(platform, contract string) (*models.FloorPrice, error) {
	parsedUrl, err := url.Parse(fmt.Sprintf(nftPriceEndpoint, platform, contract))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse platforms url")

	}
	var response models.NftResponse
	err = s.connector.Get(parsedUrl, &response)
	return &response.FloorPrice, nil
}
