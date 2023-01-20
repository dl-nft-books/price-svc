package handlers

import (
	"gitlab.com/tokend/nft-books/price-svc/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
)

func GetPlatforms(w http.ResponseWriter, r *http.Request) {
	platforms := Platforms(r)

	//hardcode Q because Q has no price
	platforms.Response.Data = append(platforms.Response.Data, resources.Platform{
		Key: resources.Key{
			ID:   "q",
			Type: resources.PLATFORMS,
		},
		Attributes: resources.PlatformAttributes{
			ChainIdentifier: 35441,
			Name:            "Q",
			Shortname:       "Q",
		},
	})
	ape.Render(w, platforms.Response)
}
