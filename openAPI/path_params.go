package openAPI

type PathParams struct {
	AbstractParameters
	mParams map[string]string
}

func NewPathParams(pParams map[string]string) *PathParams {
	zPathParams := PathParams{mParams: pParams}
	zPathParams.ParamSource = zPathParams.getKey
	return &zPathParams
}

func (this *PathParams) getKey(pKey string) string {
	return this.mParams[pKey]
}
