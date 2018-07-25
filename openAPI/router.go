/*
 * Simple To Do API
 */

package openAPI

import (
	"net/http"
	"github.com/gorilla/mux"
)

const (
	APPLICATION_JSON = "application/json; charset=UTF-8"
	TEXT_PLAIN = "text/plain; charset=UTF-8"
)

type healthCheckWrapper struct {
	mPath     string
	mFunction HealthCheckFunction
	mVersions []string
}

func (this *healthCheckWrapper) asHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		zResponse := this.mFunction(this.mVersions)
		zResponse.ApplyTo(w)
	}
}

func NewRouter(pHealthCheckPath string, pHealthCheckFunction HealthCheckFunction,
	p1stVersionedRoutes *VersionedRoutes, pAdditionalVersionedRoutes ...*VersionedRoutes) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	zWrapper := &healthCheckWrapper{mPath: pHealthCheckPath, mFunction: pHealthCheckFunction,
		mVersions: []string{p1stVersionedRoutes.GetVersion()}}
	addRoutes(router, p1stVersionedRoutes.GetRoutes())
	for _, zVersionedRoutes := range pAdditionalVersionedRoutes {
		zWrapper.mVersions = append(zWrapper.mVersions, zVersionedRoutes.GetVersion())
		addRoutes(router, zVersionedRoutes.GetRoutes())
	}
	addRoute(router, RouteGet(
		"HealthCheck",
		pHealthCheckPath,
		zWrapper.asHandler()))

	return router
}

func addRoutes(pRouter *mux.Router, pRoutes []Route) {
	for _, route := range pRoutes {
		addRoute(pRouter, route)
	}
}

func addRoute(pRouter *mux.Router, pRoute Route) {
	var handler http.Handler
	handler = pRoute.HandlerFunc
	handler = Logger(handler, pRoute.Name)

	pRouter.
		Methods(pRoute.Method).
		Path(pRoute.Pattern).
		Name(pRoute.Name).
		Handler(handler)
}

//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "HealthCheck")
//}
