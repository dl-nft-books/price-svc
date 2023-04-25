package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"time"
)

const yamlUpdatePricerKey = "update_price"

type UpdatePricer interface {
	UpdatePricerCfg() UpdatePricerCfg
}

type updatePricer struct {
	getter kv.Getter
	once   comfig.Once
}

func NewUpdatePricer(getter kv.Getter) UpdatePricer {
	return &updatePricer{
		getter: getter,
	}
}

type RunnerData struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"max_abnormal_period"`
}

type UpdatePricerCfg struct {
	Name   string     `fig:"name"`
	Runner RunnerData `fig:"runner,required"`
}

func (t *updatePricer) UpdatePricerCfg() UpdatePricerCfg {
	return t.once.Do(func() interface{} {
		var cfg UpdatePricerCfg

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(t.getter, yamlUpdatePricerKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out update pricer fields"))
		}
		return cfg
	}).(UpdatePricerCfg)
}
