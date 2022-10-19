package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval"
)

type GetPriceRequest struct {
	Platform string `url:"platform"`
	Contract string `url:"contract"`
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
		"platform=": validation.Validate(r.Platform, validation.Required),
	}.Filter()
}
