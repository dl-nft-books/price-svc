package service

import (
	"net"
	"net/http"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/tokend/nft-books/price-svc/internal/config"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko/models"
	"gitlab.com/tokend/nft-books/price-svc/resources"
)

type service struct {
	log       *logan.Entry
	copus     types.Copus
	listener  net.Listener
	coingecko *coingecko.Service
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:       cfg.Log(),
		copus:     cfg.Copus(),
		listener:  cfg.Listener(),
		coingecko: cfg.Coingecko(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

func (s *service) getCoigeckoPlatforms() (*models.Platforms, error) {
	platforms, err := s.coingecko.GetPlatforms()
	if err != nil {
		return nil, err
	}

	platformsResp := resources.PlatformListResponse{}
	mapped := make(map[string]string)
	for _, platform := range platforms {
		mapped[platform.Name] = platform.ID

		platformsResp.Data = append(platformsResp.Data, resources.Platform{
			Key: resources.Key{
				ID:   platform.ID,
				Type: resources.PLATFORMS,
			},
			Attributes: resources.PlatformAttributes{
				ChainIdentifier: int32(platform.ChainIdentifier),
				Name:            platform.Name,
				Shortname:       platform.Shortname,
			},
		})
	}

	return &models.Platforms{
		Mapped:   mapped,
		Response: platformsResp,
	}, nil
}
