package coingecko

import (
	"net/url"

	"github.com/pkg/errors"

	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko/models"
)

const (
	platformsEndpoint = "/asset_platforms"
)

func (s *Service) GetPlatforms() (models.PlatformResponse, error) {
	var response models.PlatformResponse

	parsedUrl, err := url.Parse(platformsEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse platforms url")
	}

	err = s.connector.Get(parsedUrl, &response)

	return response, err
}
