package connector

import (
	"fmt"
	"github.com/dl-nft-books/network-svc/connector/models"
	"github.com/dl-nft-books/network-svc/resources"
)

const (
	networksDetailedEndpoint = "networks/detailed"
)

func (c *Connector) GetNetworkDetailedByChainID(chainID int64) (*models.NetworkDetailedResponse, error) {
	var result resources.NetworkDetailedResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%v", c.baseUrl, networksDetailedEndpoint, chainID)
	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}
	return models.NewDetailedFromResources(result.Data)
}

func (c *Connector) GetNetworksDetailed() (*models.NetworkDetailedListResponse, error) {
	var result resources.NetworkDetailedListResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s", c.baseUrl, networksDetailedEndpoint)
	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	networks := make([]models.NetworkDetailedResponse, len(result.Data))

	for i, network := range result.Data {
		netAsModel, err := models.NewDetailedFromResources(network)
		if err != nil {
			return nil, err
		}
		networks[i] = *netAsModel
	}

	return &models.NetworkDetailedListResponse{
		Data: networks,
	}, nil
}
