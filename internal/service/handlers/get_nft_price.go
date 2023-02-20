package handlers

import (
	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko/models"
	"gitlab.com/tokend/nft-books/price-svc/resources"
	"net/http"
	"strconv"

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
	price, err := getNftPrice(r, request.Platform, coingeckoContract)
	if err != nil {
		ape.Render(w, problems.InternalError())
		Log(r).WithError(err).Error("failed to get price")
		return
	}

	if price.Usd == 0 {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response := resources.NftPriceResponse{
		Data: resources.NftPrice{
			Key: resources.Key{
				ID:   request.Contract,
				Type: resources.NFT_PRICE,
			},
			Attributes: resources.NftPriceAttributes{
				NativeCurrency: price.NativeCurrency,
				Usd:            price.Usd,
			},
		},
	}
	ape.Render(w, response)
}

func getNftPrice(r *http.Request, platform, contract string) (*models.FloorPrice, error) {
	if mockedPlatform, ok := MockedPlatforms(r)[platform]; ok {
		tokenPrice, err := strconv.ParseFloat(mockedPlatform.PricePerOneToken, 32)
		if err != nil {
			return nil, err
		}
		NftPrice, err := strconv.ParseFloat(mockedPlatform.PricePerOneNft, 32)
		if err != nil {
			return nil, err
		}
		return &models.FloorPrice{
			NativeCurrency: float32(NftPrice) / float32(tokenPrice),
			Usd:            float32(NftPrice),
		}, nil
	}
	return Coingecko(r).GetNftPrice(platform, contract)
}
