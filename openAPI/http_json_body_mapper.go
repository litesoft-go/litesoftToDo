package openAPI

import (
	"io"
	"encoding/json"
)

func HttpRequestBodyToJsonBinding(pBindingStruct interface{}, pReadCloser io.ReadCloser) (rNoBody bool, err error) {
	if pReadCloser == nil {
		rNoBody = true
		return
	}
	defer cleanUp(pReadCloser)
	err = json.NewDecoder(pReadCloser).Decode(pBindingStruct)
	return
}

func cleanUp(pReadCloser io.ReadCloser) {
	defer closeUp(pReadCloser)
	zBuf := make([]byte, 4096)
	for {
		_, err := pReadCloser.Read(zBuf)
		if err == io.EOF {
			return
		}
	}
}

func closeUp(pReadCloser io.ReadCloser) {
	pReadCloser.Close()
}
