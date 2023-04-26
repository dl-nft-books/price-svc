package service

import (
	"github.com/dl-nft-books/price-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxCoingecko(s.coingecko),
			handlers.CtxMockedTokens(s.mocked.Tokens),
			handlers.CtxMockedNfts(s.mocked.Nfts),
			handlers.CtxMockedPlatforms(s.mocked.Platforms),

			// Connectors
			handlers.CtxNetworker(s.networker),
		),
	)

	r.Route("/integrations/pricer", func(r chi.Router) {
		r.Get("/price", handlers.GetPrice)
		r.Get("/nft", handlers.GetNftPrice)
	})

	return r
}
