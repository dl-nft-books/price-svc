package responses

import "gitlab.com/tokend/nft-books/price-svc/resources"

func GetPriceResponse(price, contract string) resources.PriceResponse {
	return resources.PriceResponse{
		Data: resources.Price{
			Key: resources.Key{
				ID:   contract,
				Type: resources.PRICES,
			},
			Attributes: resources.PriceAttributes{
				Price: price,
			},
		},
	}
}
