package handlers

import (
	"context"
	networker "github.com/dl-nft-books/network-svc/connector"
	"github.com/dl-nft-books/price-svc/internal/config"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/dl-nft-books/price-svc/internal/service/coingecko"
	"github.com/dl-nft-books/price-svc/internal/service/coingecko/models"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	coingeckoCtxKey
	platformsCtxKey
	networkerCtxKey
	mockedTokensCtxKey
	mockedNftsCtxKey
	mockedPlatformsCtxKey
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

func CtxMockedNfts(mockedNfts map[string]string) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mockedNftsCtxKey, mockedNfts)
	}
}

func MockedNfts(r *http.Request) map[string]string {
	return r.Context().Value(mockedNftsCtxKey).(map[string]string)
}

func CtxMockedPlatforms(mockedPlatforms map[string]config.MockedPlatform) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mockedPlatformsCtxKey, mockedPlatforms)
	}
}

func MockedPlatforms(r *http.Request) map[string]config.MockedPlatform {
	return r.Context().Value(mockedPlatformsCtxKey).(map[string]config.MockedPlatform)
}

func CtxPlatforms(entry models.Platforms) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, platformsCtxKey, entry)
	}
}

func Platforms(r *http.Request) models.Platforms {
	return r.Context().Value(platformsCtxKey).(models.Platforms)
}
