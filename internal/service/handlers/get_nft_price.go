package handlers

import (
	"github.com/dl-nft-books/price-svc/internal/config"
	"github.com/dl-nft-books/price-svc/internal/service/requests"
	"github.com/dl-nft-books/price-svc/internal/service/responses"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetNftPrice(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPriceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	platform, ok := MockedPlatforms(r)[request.ChainId]
	if !ok {
		Log(r).WithError(err).Error("platform with such chain id doesn't exists")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	coingeckoContract := request.Contract
	if mockedToken, ok := MockedNfts(r)[request.Contract]; ok {
		coingeckoContract = mockedToken
	}
	price, err := getNftPrice(r, platform, coingeckoContract)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		Log(r).WithError(err).Error("failed to get price")
		return
	}

	if price == "" {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, responses.GetNftPriceResponse(price, request.Contract))
}

func getNftPrice(r *http.Request, platform config.MockedPlatform, contract string) (string, error) {
	if cast.ToFloat64(platform.PricePerOneNft) > 0 {
		return platform.PricePerOneNft, nil
	}
	return Coingecko(r).GetNftPrice(platform.Id, contract)
}
