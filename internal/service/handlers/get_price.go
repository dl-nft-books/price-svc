package handlers

import (
	"fmt"
	"github.com/dl-nft-books/price-svc/internal/data"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"

	"github.com/dl-nft-books/price-svc/internal/service/requests"
	"github.com/dl-nft-books/price-svc/internal/service/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPrice(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPriceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	coingeckoContract := request.Contract
	if mockedToken, ok := MockedTokens(r)[request.Contract]; ok {
		coingeckoContract = mockedToken
	}

	price, err := getPrice(r, request.Platform, coingeckoContract)
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
		key = request.Platform
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
func getPrice(r *http.Request, platform, contract string) (string, error) {
	if mockedPlatforms, ok := MockedPlatforms(r)[platform]; ok {
		if cast.ToFloat64(mockedPlatforms.PricePerOneToken) > 0 {
			return mockedPlatforms.PricePerOneToken, nil
		}
	}
	fmt.Println("didn't work :(")
	fmt.Println(MockedPlatforms(r))
	return Coingecko(r).GetPrice(platform, contract, "usd")
}
