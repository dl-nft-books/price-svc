package handlers

import (
	"github.com/dl-nft-books/price-svc/internal/config"
	"github.com/dl-nft-books/price-svc/internal/service/eth_reader"
	"github.com/dl-nft-books/price-svc/internal/service/requests"
	"github.com/dl-nft-books/price-svc/internal/service/responses"
	"github.com/ethereum/go-ethereum/common"
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
	ethReader := eth_reader.NewEthReader(networker)

	erc20Data, err := ethReader.GetErc20Data(common.HexToAddress(request.Contract))
	if err != nil {
		Log(r).Error(errors.Wrap(err, "failed to get erc20 data"))
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.GetPriceResponse(price, key, *erc20Data))
}
func getPrice(r *http.Request, platform config.MockedPlatform, contract string) (string, error) {
	if cast.ToFloat64(platform.PricePerOneToken) > 0 {
		return platform.PricePerOneToken, nil
	}
	return Coingecko(r).GetPrice(platform.Id, contract, "usd")
}
