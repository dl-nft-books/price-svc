package responses

import (
	"github.com/dl-nft-books/price-svc/internal/data"
	"github.com/dl-nft-books/price-svc/resources"
)

func GetPriceResponse(price string, contract string, erc20Data data.Erc20Data) resources.PriceResponse {
	return resources.PriceResponse{
		Data: resources.Price{
			Key: resources.Key{
				ID:   contract,
				Type: resources.PRICES,
			},
			Attributes: resources.PriceAttributes{
				Price: price,
				Token: erc20Data.Resource(),
			},
		},
	}
}
