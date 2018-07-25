package openAPI

import "net/http"

type Response struct {
	mCode        int
	mContentType string
	mText        string
}

func NewResponse(pCode int, pContentType string, pText string) *Response {
	return &Response{mCode: pCode, mContentType: pContentType, mText: pText}
}

func (this *Response) GetCode() int {
	return this.mCode
}

func (this *Response) GetContentType() string {
	return this.mContentType
}

func (this *Response) GetText() string {
	return this.mText
}

func (this *Response) ApplyTo(w http.ResponseWriter) {
	w.Header().Set("Content-Type", this.mContentType)
	w.WriteHeader(this.mCode)
	w.Write([]byte(this.mText))
}

func NewResponseOK(pContentType string, pText string) *Response {
	return NewResponse(http.StatusOK, pContentType, pText)
}

func NewTextResponseOK(pText string) *Response {
	return NewResponseOK(TEXT_PLAIN, pText)
}

func NewJsonResponseOK(pText string) *Response {
	return NewResponseOK(APPLICATION_JSON, pText)
}

func NewJsonResponseCreated(pText string) *Response {
	return NewResponse(http.StatusCreated, APPLICATION_JSON, pText)
}

func NewTextResponseCreated(pWhat string) *Response {
	return NewResponse(http.StatusCreated, TEXT_PLAIN, pWhat + " created")
}

func NewTextResponseUpdated(pWhat string) *Response {
	return NewResponse(http.StatusCreated, TEXT_PLAIN, pWhat + " updated")
}

func NewInvalidTextResponse(pMessage string) *Response {
	return NewResponse(http.StatusBadRequest, TEXT_PLAIN, pMessage)
}

func NewNotFoundTextResponse(pMessage string) *Response {
	return NewResponse(http.StatusNotFound, TEXT_PLAIN, pMessage)
}

func NewFailedTextResponse(err error) *Response {
	return NewResponse(http.StatusInternalServerError, TEXT_PLAIN, err.Error())
}

type HealthCheckFunction func(pVersions []string) *Response
