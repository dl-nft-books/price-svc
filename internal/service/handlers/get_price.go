package handlers

import (
	"fmt"
	"github.com/dl-nft-books/price-svc/internal/config"
	"github.com/dl-nft-books/price-svc/internal/data"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/dl-nft-books/price-svc/internal/service/requests"
	"github.com/dl-nft-books/price-svc/internal/service/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

var (
	CachePlatforms map[string]config.MockedPlatform
	lastUpdated    map[string]time.Time
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
	if len(CachePlatforms) == 0 {
		CachePlatforms = make(map[string]config.MockedPlatform)
	}
	if len(lastUpdated) == 0 {
		lastUpdated = make(map[string]time.Time)
	}
	if mockedPlatforms, ok := CachePlatforms[platform]; ok &&
		time.Now().Sub(lastUpdated[platform]).Minutes() < 5 {
		return mockedPlatforms.PricePerOneToken, nil
	}
	if platform == "q" {
		CachePlatforms[platform] = config.MockedPlatform{
			Id:               platform,
			ChainId:          MockedPlatforms(r)[platform].ChainId,
			Name:             MockedPlatforms(r)[platform].Name,
			ShortName:        MockedPlatforms(r)[platform].ShortName,
			PricePerOneToken: MockedPlatforms(r)[platform].PricePerOneToken,
			PricePerOneNft:   MockedPlatforms(r)[platform].PricePerOneNft,
		}
		return MockedPlatforms(r)[platform].PricePerOneToken, nil
	}
	price, err := Coingecko(r).GetPrice(platform, contract, "usd")
	if err != nil {
		return "", errors.Wrap(err, "failed to get price")
	}
	if len(CachePlatforms) == 0 {
		CachePlatforms = make(map[string]config.MockedPlatform)
	}
	CachePlatforms[platform] = config.MockedPlatform{
		Id:               platform,
		ChainId:          MockedPlatforms(r)[platform].ChainId,
		Name:             MockedPlatforms(r)[platform].Name,
		ShortName:        MockedPlatforms(r)[platform].ShortName,
		PricePerOneToken: price,
		PricePerOneNft:   MockedPlatforms(r)[platform].PricePerOneNft,
	}
	lastUpdated[platform] = time.Now()
	Log(r).Debug(fmt.Sprintf("Set new price for %v in %v", platform, lastUpdated[platform].Format("02 Jan 2006 3:04:05 pm")))
	fmt.Println("CachePlatforms", CachePlatforms)
	return price, nil
}
