package connector

import (
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/signed"
)

type Pricer interface {
	PricerConnector() *Connector
}

type pricer struct {
	confOnce comfig.Once

	getter kv.Getter
}

func NewPricer(getter kv.Getter) Pricer {
	return &pricer{
		getter: getter,
	}
}

type pricerConfig struct {
	URL *url.URL `fig:"url,required"`
}

func (c *pricer) PricerConnector() *Connector {
	return c.confOnce.Do(func() interface{} {
		var config pricerConfig

		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "connector")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out pricer"))
		}

		cli := signed.NewClient(http.DefaultClient, config.URL)

		return NewConnector(cli)
	}).(*Connector)
}
