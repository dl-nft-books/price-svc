package config

import (
	"gitlab.com/distributed_lab/kit/kv"
)

const mockedTokensYamlKey = "mocked_tokens"

func (c *config) MockedTokens() map[string]string {
	return c.mockedTokensOnce.Do(func() interface{} {
		return kv.MustGetStringMap(c.getter, mockedTokensYamlKey)
	}).(map[string]string)
}
