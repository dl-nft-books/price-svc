package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

const ethererYamlKey = "ethereum_client"

type EtherClient struct {
	Rpc       *ethclient.Client `fig:"rpc,required"`
	WebSocket *ethclient.Client `fig:"web_socket,required"`
}

func (c *config) EtherClient() EtherClient {
	return c.ethererOnce.Do(func() interface{} {
		var cfg EtherClient

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks, ethClientHook).
			From(kv.MustGetStringMap(c.getter, ethererYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint tracker config"))
		}

		return cfg
	}).(EtherClient)
}

var ethClientHook = figure.Hooks{
	"*ethclient.Client": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			client, err := ethclient.Dial(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to convert value into ethclient")
			}
			return reflect.ValueOf(client), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
