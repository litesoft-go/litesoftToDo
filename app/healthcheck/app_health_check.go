package healthcheck

import (
	"encoding/json"
	"litesoftToDo/app/persistence"
	"litesoftToDo/openAPI"
)

type HealthyBody struct {
	Lists uint `json:"lists"`

	Versions []string `json:"supportedVersions"`
}

type AppHealthChecker struct {
	mPersistor persistence.Persistor
}

func New(pPersistor persistence.Persistor) *AppHealthChecker {
	return &AppHealthChecker{mPersistor: pPersistor}
}

func (this *AppHealthChecker) AsHealthCheckFunction() openAPI.HealthCheckFunction {
	return func(pVersions []string) *openAPI.Response {
		zCount, err := this.mPersistor.ListCount()
		if err != nil {
			return openAPI.NewFailedTextResponse(err)
		}
		zBody := HealthyBody{Lists: zCount, Versions: pVersions}
		zBytes, err := json.Marshal(zBody)
		if err != nil {
			return openAPI.NewFailedTextResponse(err)
		}
		return openAPI.NewJsonResponseOK(string(zBytes))
	}
}
