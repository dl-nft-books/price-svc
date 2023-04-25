package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"

	networker "github.com/dl-nft-books/network-svc/connector"
	"github.com/dl-nft-books/price-svc/internal/service/coingecko"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	Coingecko() *coingecko.Service
	Mocked() MockedStructures
	UpdatePricer

	networker.NetworkConfigurator
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	getter           kv.Getter
	ethererOnce      comfig.Once
	coingecko        comfig.Once
	mockedTokensOnce comfig.Once
	UpdatePricer
	networker.NetworkConfigurator
}

func New(getter kv.Getter) Config {
	return &config{
		getter:              getter,
		Copuser:             copus.NewCopuser(getter),
		Listenerer:          comfig.NewListenerer(getter),
		Logger:              comfig.NewLogger(getter, comfig.LoggerOpts{}),
		NetworkConfigurator: networker.NewNetworkConfigurator(getter),
		UpdatePricer:        NewUpdatePricer(getter),
	}
}
