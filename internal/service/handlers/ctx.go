package handlers

import (
	"context"
	networker "gitlab.com/tokend/nft-books/network-svc/connector"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko"
	"gitlab.com/tokend/nft-books/price-svc/internal/service/coingecko/models"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	coingeckoCtxKey
	platformsCtxKey
	ethReaderCtxKey
	networkerCtxKey
	mockedTokensCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxCoingecko(entry *coingecko.Service) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, coingeckoCtxKey, entry)
	}
}

func Coingecko(r *http.Request) *coingecko.Service {
	return r.Context().Value(coingeckoCtxKey).(*coingecko.Service)
}

func CtxNetworker(entry *networker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networkerCtxKey, entry)
	}
}

func Networker(r *http.Request) *networker.Connector {
	return r.Context().Value(networkerCtxKey).(*networker.Connector)
}

func CtxMockedTokens(mockedTokens map[string]string) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mockedTokensCtxKey, mockedTokens)
	}
}

func MockedTokens(r *http.Request) map[string]string {
	return r.Context().Value(mockedTokensCtxKey).(map[string]string)
}

func CtxPlatforms(entry *models.Platforms) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, platformsCtxKey, entry)
	}
}

func Platforms(r *http.Request) *models.Platforms {
	return r.Context().Value(platformsCtxKey).(*models.Platforms)
}
