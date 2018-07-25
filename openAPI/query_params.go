package openAPI

import (
	"net/url"
)

type QueryParams struct {
	AbstractParameters
	mParams url.Values
}

func NewQueryParams(pParams url.Values) *QueryParams {
	zQueryParams := QueryParams{mParams: pParams}
	zQueryParams.ParamSource = zQueryParams.getKey
	return &zQueryParams
}

func (this *QueryParams) getKey(pKey string) string {
	return this.mParams.Get(pKey)
}
