/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NetworkDetailedAttributes struct {
	// Chain id
	ChainId int64 `json:"chain_id"`
	// decimal
	Decimals int64 `json:"decimals"`
	// Address of token factory on current network
	FactoryAddress string `json:"factory_address"`
	// Token factory name
	FactoryName string `json:"factory_name"`
	// Token factory version
	FactoryVersion string `json:"factory_version"`
	// A first block to begin monitoring events
	FirstBlock int64 `json:"first_block"`
	// Network name
	Name string `json:"name"`
	// RPC url to listen events
	RpcUrl string `json:"rpc_url"`
	// Native token name
	TokenName string `json:"token_name"`
	// Native token symbol
	TokenSymbol string `json:"token_symbol"`
	// Websocket url to listen events
	WsUrl string `json:"ws_url"`
}
