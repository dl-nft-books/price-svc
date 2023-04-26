package models

type FloorPrice struct {
	Usd float64 `json:"usd"`
}

type NftResponse struct {
	FloorPrice FloorPrice `json:"floor_price"`
	Error      string     `json:"error"`
}
