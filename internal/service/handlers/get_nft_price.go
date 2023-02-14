package handlers

import (
	"gitlab.com/tokend/nft-books/price-svc/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/requests"
)

func GetNftPrice(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPriceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	coingeckoContract := request.Contract
	if mockedToken, ok := MockedNfts(r)[request.Contract]; ok {
		coingeckoContract = mockedToken
	}
	price, err := Coingecko(r).GetNftPrice(request.Platform, coingeckoContract)
	if err != nil {
		ape.Render(w, problems.InternalError())
		Log(r).WithError(err).Error("failed to get price")
		return
	}

	if price == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response := resources.NftPrice{
		Key: resources.Key{
			ID:   request.Contract,
			Type: resources.NFT_PRICE,
		},
		Attributes: resources.NftPriceAttributes{
			NativeCurrency: price.NativeCurrency,
			Usd:            price.Usd,
		},
	}
	ape.Render(w, response)
}
