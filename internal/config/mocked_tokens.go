package config

import (
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

const mockedYamlKey = "mocked"

type MockedToken struct {
	ActualAddress    string `fig:"actual_address,required"`
	CoingeckoAddress string `fig:"coingecko_address,required"`
}
type MockedPlatform struct {
	Id               string `fig:"id,required"`
	ChainId          int32  `fig:"chain_id,required"`
	Name             string `fig:"name,required"`
	ShortName        string `fig:"short_name,required"`
	PricePerOneToken string `fig:"price_per_one_token"`
	PricePerOneNft   string `fig:"price_per_one_nft"`
}

type MockedStructures struct {
	Tokens    map[string]string
	Nfts      map[string]string
	Platforms map[string]MockedPlatform
}

func (c *config) Mocked() MockedStructures {
	return c.mockedTokensOnce.Do(func() interface{} {
		cfg := struct {
			MockedTokens    []MockedToken    `fig:"tokens"`
			MockedNfts      []MockedToken    `fig:"nfts"`
			MockedPlatforms []MockedPlatform `fig:"platforms"`
		}{}
		result := MockedStructures{
			Tokens:    make(map[string]string),
			Nfts:      make(map[string]string),
			Platforms: make(map[string]MockedPlatform),
		}

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks, mockedHook).
			From(kv.MustGetStringMap(c.getter, mockedYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mocked tokens map"))
		}

		for _, mockedToken := range cfg.MockedTokens {
			result.Tokens[mockedToken.ActualAddress] = mockedToken.CoingeckoAddress
		}
		for _, mockedNft := range cfg.MockedNfts {
			result.Nfts[mockedNft.ActualAddress] = mockedNft.CoingeckoAddress
		}
		for _, mockedPlatform := range cfg.MockedPlatforms {
			result.Platforms[mockedPlatform.Id] = mockedPlatform
		}
		return result
	}).(MockedStructures)
}

var mockedHook = figure.Hooks{
	"[]config.MockedToken": func(value interface{}) (reflect.Value, error) {
		switch s := value.(type) {
		case []interface{}:
			mockedTokens := make([]MockedToken, 0)
			for _, rawElement := range s {
				mapElement, ok := rawElement.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.New("failed to cast mapElement to interface")
				}

				normalizedMap := make(map[string]interface{}, len(mapElement))
				for key, value := range mapElement {
					keyAsString := fmt.Sprintf("%v", key)
					normalizedMap[keyAsString] = value
				}

				var mockedToken MockedToken
				if err := figure.
					Out(&mockedToken).
					With(figure.BaseHooks).
					From(normalizedMap).
					Please(); err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out mockedToken from normalized map")
				}

				mockedTokens = append(mockedTokens, mockedToken)
			}

			return reflect.ValueOf(mockedTokens), nil
		default:
			return reflect.Value{}, errors.New("unexpected type while figuring out []MockedToken")
		}
	},
	"[]config.MockedPlatform": func(value interface{}) (reflect.Value, error) {
		switch s := value.(type) {
		case []interface{}:
			mockedPlatforms := make([]MockedPlatform, 0)
			for _, rawElement := range s {
				mapElement, ok := rawElement.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.New("failed to cast mapElement to interface")
				}

				normalizedMap := make(map[string]interface{}, len(mapElement))
				for key, value := range mapElement {
					keyAsString := fmt.Sprintf("%v", key)
					normalizedMap[keyAsString] = value
				}

				var mockedPlatform MockedPlatform
				if err := figure.
					Out(&mockedPlatform).
					With(figure.BaseHooks).
					From(normalizedMap).
					Please(); err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out mockedPlatform from normalized map")
				}

				mockedPlatforms = append(mockedPlatforms, mockedPlatform)
			}

			return reflect.ValueOf(mockedPlatforms), nil
		default:
			return reflect.Value{}, errors.New("unexpected type while figuring out []MockedPlatform")
		}
	},
	"*int64": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToInt64E(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse *int64")
		}
		return reflect.ValueOf(&result), nil
	},
}
