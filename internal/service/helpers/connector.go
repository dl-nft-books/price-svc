package helpers

import (
	"context"
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

func (c *Connector) PostJSON(endpoint *url.URL, req interface{}, ctx context.Context, dst interface{}) (err error) {
	_, response, err := c.base.PostJSON(endpoint, req, ctx)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
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
