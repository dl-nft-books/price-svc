package connector

import (
	"encoding/json"
	"net/url"

	"gitlab.com/distributed_lab/json-api-connector/base"
	"gitlab.com/distributed_lab/json-api-connector/client"
)

type Connector struct {
	base *base.Connector
}

func NewConnector(client client.Client) *Connector {
	return &Connector{
		base: base.NewConnector(client),
	}
}

func (c *Connector) Get(endpoint *url.URL, dst interface{}) (err error) {
	response, err := c.base.Get(endpoint)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}
