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

	cg := Coingecko(r)
	price, err := cg.GetPrice(request.Platform, request.Contract, "usd")
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

	ape.Render(w, responses.GetPriceResponse(price, key))
}
