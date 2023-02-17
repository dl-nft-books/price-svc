package data

import (
	"gitlab.com/tokend/nft-books/price-svc/resources"
)

type Erc20Data struct {
	Symbol   string
	Name     string
	Decimals int32
}

func (erc20 *Erc20Data) Resource() resources.Token {
	return resources.Token{
		Name:     erc20.Name,
		Symbol:   erc20.Symbol,
		Decimals: erc20.Decimals,
	}
}
