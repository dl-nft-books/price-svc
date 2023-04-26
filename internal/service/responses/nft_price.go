package responses

import (
	"github.com/dl-nft-books/price-svc/resources"
)

func GetNftPriceResponse(price string, contract string) resources.NftPriceResponse {
	return resources.NftPriceResponse{
		Data: resources.NftPrice{
			Key: resources.Key{
				ID:   contract,
				Type: resources.PRICES,
			},
			Attributes: resources.NftPriceAttributes{
				FloorPrice: price,
			},
		},
	}
}
