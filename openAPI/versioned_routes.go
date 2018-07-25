package openAPI

import "net/http"

type VersionedRoutes struct {
	mPrefix string
	mVersion string
	mRoutes  []Route
}

func NewVersionedRoutes(pPrefix, pVersion string) *VersionedRoutes {
	return &VersionedRoutes{mPrefix:pPrefix, mVersion: pVersion, mRoutes: []Route{}}
}

func (this *VersionedRoutes) GetVersion() string {
	return this.mVersion
}

func (this *VersionedRoutes) GetRoutes() []Route {
	return this.mRoutes
}

func (this *VersionedRoutes) AddGet(pName string, pHandlerFunc http.HandlerFunc, pPostVersionPattern string) *VersionedRoutes {
	return this.addRoute(RouteGet(pName, this.mPrefix + this.mVersion + pPostVersionPattern, pHandlerFunc))
}

func (this *VersionedRoutes) AddRoutePost(pName string, pHandlerFunc http.HandlerFunc, pPostVersionPattern string) *VersionedRoutes {
	return this.addRoute(RoutePost(pName, this.mPrefix + this.mVersion + pPostVersionPattern, pHandlerFunc))
}

func (this *VersionedRoutes) addRoute(pRoute Route) *VersionedRoutes {
	this.mRoutes = append(this.mRoutes, pRoute)
	return this
}
