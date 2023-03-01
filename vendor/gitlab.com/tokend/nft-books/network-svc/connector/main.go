package connector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Connector struct {
	token   string
	baseUrl string
}

func NewConnector(authToken, serviceUrl string) *Connector {
	return &Connector{authToken, serviceUrl}
}

func (c *Connector) get(endpoint string, dst interface{}) error {
	// creating request
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create connector request")
	}

	//  setting headers
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// sending request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to process request")
	}
	if response == nil {
		return errors.New("failed to process request: response is nil")
	}

	defer response.Body.Close()

	// parsing response
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	return json.Unmarshal(raw, &dst)
}
