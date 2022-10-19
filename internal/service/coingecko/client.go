package coingecko

import (
	"net/http"
	"net/url"
	"path"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Client struct {
	client *http.Client
	base   *url.URL
	apiKey string
}

func NewClient(client *http.Client, base *url.URL, apiKey string) *Client {
	return &Client{
		client: client,
		base:   base,
		apiKey: apiKey,
	}
}

func (c *Client) Resolve(endpoint *url.URL) (string, error) {
	q := endpoint.Query()
	q.Add("x_cg_pro_api_key", c.apiKey)
	endpoint.RawQuery = q.Encode()

	u := *c.base
	basePath := u.Path
	prevPath := endpoint.Path

	if basePath != "" {
		endpoint.Path = path.Join(basePath, endpoint.Path)
		u.Path = ""
	}

	resolved := u.ResolveReference(endpoint)
	endpoint.Path = prevPath

	return resolved.String(), nil
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	r.Header.Set("Content-Type", "application/json")

	response, err := c.client.Do(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform http request")
	}

	return response, nil
}
