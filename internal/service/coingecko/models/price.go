package models

import (
	"fmt"
	"strings"
)

type PriceResponse map[string]map[string]float64

func (p PriceResponse) GetPrice(contract, vs string) string {
	data, ok := p[strings.ToLower(contract)]
	if !ok {
		return ""
	}

	price, ok := data[strings.ToLower(vs)]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%f", price)
}
