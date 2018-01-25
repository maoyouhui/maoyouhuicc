package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

var RouteMap = []*rest.Route{
	//auth
	rest.Post("/wxconfig", getWxConfig),
}
