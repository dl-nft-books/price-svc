package handlers

import (
	"gitlab.com/tokend/nft-books/price-svc/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
)

func GetPlatforms(w http.ResponseWriter, r *http.Request) {
	platforms := Platforms(r)

	//hardcode Q because Q has no price
	for _, mockedPlatform := range MockedPlatforms(r) {
		platforms.Response.Data = append(platforms.Response.Data, resources.Platform{
			Key: resources.Key{
				ID:   mockedPlatform.Id,
				Type: resources.PLATFORMS,
			},
			Attributes: resources.PlatformAttributes{
				ChainIdentifier: mockedPlatform.ChainId,
				Name:            mockedPlatform.Name,
				Shortname:       mockedPlatform.ShortName,
			},
		})
	}

	ape.Render(w, platforms.Response)
}
