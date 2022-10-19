package config

import (
	"net/http"
	"net/url"
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"

	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/helpers"
)

func (c *config) Coingecko() *coingecko.Service {
	return c.coingecko.Do(func() interface{} {
		var config struct {
			Url        *url.URL      `fig:"url,required"`
			ApiKey     string        `fig:"api_key"`
			Expiration time.Duration `fig:"expiration"`
		}

		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "coingecko")).
			Please()
		if err != nil {
			panic(err)
		}

		cli := coingecko.NewClient(http.DefaultClient, config.Url, config.ApiKey)

		return coingecko.NewService(helpers.NewConnector(cli), config.Expiration)
	}).(*coingecko.Service)
}
