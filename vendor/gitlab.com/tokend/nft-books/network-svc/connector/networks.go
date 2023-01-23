package connector

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
	"gitlab.com/tokend/nft-books/network-svc/resources"
	"log"
)

const (
	networksEndpoint         = "networks"
	networksDetailedEndpoint = "networks/detailed"
)

func (c *Connector) GetNetworkDetailedByChainID(chainID int64) (*models.NetworkDetailedResponse, error) {
	var result resources.NetworkDetailedResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%v", c.baseUrl, networksDetailedEndpoint, chainID)
	log.Println("NETWORK CHAIN ID", chainID)
	log.Println("NETWORK CHAIN FULL ENDPOINT", fullEndpoint)
	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		log.Println("NET ERR", err)
		return nil, err
	}
	log.Println("NETWORK RES", result)

	rpc, err := ethclient.Dial(result.Data.Attributes.RpcUrl)
	if err != nil {
		return nil, err
	}
	ws, err := ethclient.Dial(result.Data.Attributes.WsUrl)
	if err != nil {
		return nil, err
	}
	return &models.NetworkDetailedResponse{
		Name:           result.Data.Attributes.Name,
		ChainId:        result.Data.Attributes.ChainId,
		RpcUrl:         rpc,
		WsUrl:          ws,
		FactoryAddress: result.Data.Attributes.FactoryAddress,
		FactoryName:    result.Data.Attributes.FactoryName,
		FactoryVersion: result.Data.Attributes.FactoryVersion,
		FirstBlock:     result.Data.Attributes.FirstBlock,
		TokenName:      result.Data.Attributes.TokenName,
		TokenSymbol:    result.Data.Attributes.TokenSymbol,
	}, nil
}

func (c *Connector) GetNetworksDetailed() (*models.NetworkDetailedListResponse, error) {
	var result resources.NetworkDetailedListResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s", c.baseUrl, networksDetailedEndpoint)
	log.Println("NETWORK CHAIN FULL ENDPOINT", fullEndpoint)
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
