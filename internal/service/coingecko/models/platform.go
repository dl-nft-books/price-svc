package models

import "github.com/dl-nft-books/price-svc/resources"

type Platform struct {
	ID              string `json:"id"`
	ChainIdentifier uint64 `json:"chain_identifier"`
	Name            string `json:"name"`
	Shortname       string `json:"shortname"`
}

type PlatformResponse []Platform

type Platforms struct {
	Mapped   map[string]string
	Response resources.PlatformListResponse
}
