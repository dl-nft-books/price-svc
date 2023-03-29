package eth_reader

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/dl-nft-books/price-svc/internal/data"
	"github.com/dl-nft-books/price-svc/solidity/generated/erc20"
)

var (
	nullAddress  = common.Address{}
	ethereumData = data.Erc20Data{
		Symbol:   "ETH",
		Name:     "Ethereum",
		Decimals: 18,
	}
)

type EthReader struct {
	rpc *ethclient.Client
}

func NewEthReader(rpc *ethclient.Client) *EthReader {
	return &EthReader{
		rpc: rpc,
	}
}

func (r *EthReader) GetErc20Data(address common.Address) (*data.Erc20Data, error) {
	if address == nullAddress {
		return &ethereumData, nil
	}

	erc20Instance, err := erc20.NewErc20(address, r.rpc)
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
