package models

type FloorPrice struct {
	NativeCurrency float32 `json:"native_currency"`
	Usd            float32 `json:"usd"`
}

type NftResponse struct {
	FloorPrice FloorPrice `json:"floor_price"`
	Error      string     `json:"error"`
}
