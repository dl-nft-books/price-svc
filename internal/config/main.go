package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"

	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	Coingecko() *coingecko.Service
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	getter kv.Getter

	coingecko comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
