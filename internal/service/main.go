package service

import (
	networker "github.com/dl-nft-books/network-svc/connector"
	"net"
	"net/http"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/dl-nft-books/price-svc/internal/config"
	"github.com/dl-nft-books/price-svc/internal/service/coingecko"
)

type service struct {
	log       *logan.Entry
	copus     types.Copus
	listener  net.Listener
	coingecko *coingecko.Service
	mocked    config.MockedStructures
	networker *networker.Connector
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
		mocked:    cfg.Mocked(),
		networker: cfg.NetworkConnector(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
