package eth_reader

import (
	"github.com/dl-nft-books/network-svc/connector/models"
	"github.com/dl-nft-books/price-svc/internal/data"
	"github.com/dl-nft-books/price-svc/solidity/generated/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	nullAddress = common.Address{}
)

type EthReader struct {
	network *models.NetworkDetailedResponse
}

func NewEthReader(network *models.NetworkDetailedResponse) *EthReader {
	return &EthReader{
		network: network,
	}
}

func (r *EthReader) GetErc20Data(address common.Address) (*data.Erc20Data, error) {
	if address == nullAddress {
		return &data.Erc20Data{
			Symbol:   r.network.TokenSymbol,
			Name:     r.network.Name,
			Decimals: int32(r.network.Decimals),
		}, nil
	}

	erc20Instance, err := erc20.NewErc20(address, r.network.RpcUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a new erc20 instance")
	}

	name, err := erc20Instance.Name(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get erc20 name")
	}

	symbol, err := erc20Instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contract's symbol")
	}

	decimals, err := erc20Instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contract's decimals")
	}

	return &data.Erc20Data{
		Symbol:   symbol,
		Name:     name,
		Decimals: int32(decimals),
	}, nil
}
