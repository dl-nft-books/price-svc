package handlers

import (
	"github.com/dl-nft-books/price-svc/internal/config"
	"github.com/dl-nft-books/price-svc/internal/data"
	"github.com/dl-nft-books/price-svc/internal/service/requests"
	"github.com/dl-nft-books/price-svc/internal/service/responses"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetPrice(w http.ResponseWriter, r *http.Request) {
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
	if mockedToken, ok := MockedTokens(r)[request.Contract]; ok {
		coingeckoContract = mockedToken
	}

	price, err := getPrice(r, platform, coingeckoContract)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		Log(r).WithError(err).Error("failed to get price")
		return
	}
	if price == "" {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	key := request.Contract
	if key == "" {
		key = platform.Id
	}

	networker, err := Networker(r).GetNetworkDetailedByChainID(request.ChainId)
	if err != nil {
		Log(r).Error(errors.Wrap(err, "failed to select network from the database"))
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, responses.GetPriceResponse(price, key, data.Erc20Data{
		Name:     networker.TokenName,
		Symbol:   networker.TokenSymbol,
		Decimals: int32(networker.Decimals),
	}))
}
func getPrice(r *http.Request, platform config.MockedPlatform, contract string) (string, error) {
	if cast.ToFloat64(platform.PricePerOneToken) > 0 {
		return platform.PricePerOneToken, nil
	}
	return Coingecko(r).GetPrice(platform.Id, contract, "usd")
}
