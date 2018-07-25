package openAPI

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

const (
	GET  = "GET"
	POST = "POST"
)

func RouteGet(pName string, pPattern string, pHandlerFunc http.HandlerFunc) Route {
	return Route{Name: pName, Method: GET, Pattern: pPattern, HandlerFunc: pHandlerFunc}
}

func RoutePost(pName string, pPattern string, pHandlerFunc http.HandlerFunc) Route {
	return Route{Name: pName, Method: POST, Pattern: pPattern, HandlerFunc: pHandlerFunc}
}
