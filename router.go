package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

var RouteMap = []*rest.Route{
	rest.Get("/wxconfig", getWxConfig),
	rest.Post("/wxconfig", getWxConfig),
}
