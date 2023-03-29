package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"github.com/dl-nft-books/price-svc/internal/service/handlers"
	"github.com/dl-nft-books/price-svc/resources"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	platforms, err := s.getCoigeckoPlatforms()
	if err != nil {
		panic(err)
	}

	//hardcode Q because Q has no price
	for _, mockedPlatform := range s.mocked.Platforms {
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
	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxCoingecko(s.coingecko),
			handlers.CtxPlatforms(platforms),
			handlers.CtxMockedTokens(s.mocked.Tokens),
			handlers.CtxMockedNfts(s.mocked.Nfts),
			handlers.CtxMockedPlatforms(s.mocked.Platforms),

			// Connectors
			handlers.CtxNetworker(s.networker),
		),
	)

	r.Route("/integrations/pricer", func(r chi.Router) {
		r.Get("/platforms", handlers.GetPlatforms)
		r.Get("/price", handlers.GetPrice)
		r.Get("/nft", handlers.GetNftPrice)
	})

	return r
}
