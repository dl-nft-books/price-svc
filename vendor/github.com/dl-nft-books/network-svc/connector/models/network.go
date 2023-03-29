package models

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/dl-nft-books/network-svc/resources"
)

type NetworkDetailedResponse struct {
	ChainId        int64             `json:"chain_id"`
	FactoryAddress string            `json:"factory_address"`
	FactoryName    string            `json:"factory_name"`
	FactoryVersion string            `json:"factory_version"`
	FirstBlock     int64             `json:"first_block"`
	Name           string            `json:"name"`
	RpcUrl         *ethclient.Client `json:"rpc_url"`
	TokenName      string            `json:"token_name"`
	TokenSymbol    string            `json:"token_symbol"`
	Decimals       int64             `json:"decimals"`
	WsUrl          *ethclient.Client `json:"ws_url"`
}

type NetworkResponse struct {
	ChainId        int64  `json:"chain_id"`
	FactoryAddress string `json:"factory_address"`
	Name           string `json:"name"`
	TokenName      string `json:"token_name"`
	TokenSymbol    string `json:"token_symbol"`
	Decimals       int64  `json:"decimals"`
}
type NetworkListResponse struct {
	Data []NetworkResponse `json:"data"`
}
type NetworkDetailedListResponse struct {
	Data []NetworkDetailedResponse `json:"data"`
}

func NewDetailedFromResources(n resources.NetworkDetailed) (*NetworkDetailedResponse, error) {
	rpc, err := ethclient.Dial(n.Attributes.RpcUrl)
	if err != nil {
		return nil, err
	}
	ws, err := ethclient.Dial(n.Attributes.WsUrl)
	if err != nil {
		return nil, err
	}
	return &NetworkDetailedResponse{
		Name:           n.Attributes.Name,
		ChainId:        n.Attributes.ChainId,
		RpcUrl:         rpc,
		WsUrl:          ws,
		FactoryAddress: n.Attributes.FactoryAddress,
		FactoryName:    n.Attributes.FactoryName,
		FactoryVersion: n.Attributes.FactoryVersion,
		FirstBlock:     n.Attributes.FirstBlock,
		TokenName:      n.Attributes.TokenName,
		TokenSymbol:    n.Attributes.TokenSymbol,
		Decimals:       n.Attributes.Decimals,
	}, nil
}
