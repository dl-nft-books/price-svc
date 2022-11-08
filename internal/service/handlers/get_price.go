package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/tokend/nft-books/price-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/responses"
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

	price, err := Coingecko(r).GetPrice(request.Platform, coingeckoContract, "usd")
	if err != nil {
		ape.Render(w, problems.InternalError())
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

	erc20Data, err := EthReader(r).GetErc20Data(common.HexToAddress(request.Contract))
	if err != nil {
		ape.Render(w, problems.InternalError())
		Log(r).WithError(err).Error("failed to get erc20 from the contract")
		return
	}

	ape.Render(w, responses.GetPriceResponse(price, key, *erc20Data))
}
