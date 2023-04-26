package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval"
)

type GetPriceRequest struct {
	Contract string `url:"contract"`
	ChainId  int64  `url:"chain_id"`
}

func NewGetPriceRequest(r *http.Request) (GetPriceRequest, error) {
	var result GetPriceRequest

	if err := urlval.Decode(r.URL.Query(), &result); err != nil {
		return result, err
	}

	return result, result.Validate()
}

func (r *GetPriceRequest) Validate() error {
	return validation.Errors{
		"chain_id=": validation.Validate(r.ChainId, validation.Required),
	}.Filter()
}
