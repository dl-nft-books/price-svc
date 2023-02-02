package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	platforms, err := s.getCoigeckoPlatforms()
	if err != nil {
		panic(err)
	}

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxCoingecko(s.coingecko),
			handlers.CtxPlatforms(platforms),
			handlers.CtxMockedTokens(s.mocked.Tokens),

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
