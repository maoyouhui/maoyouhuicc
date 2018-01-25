package main

import "github.com/ant0ine/go-json-rest/rest"
import (
	"net/http"
	"log"
)

func main()  {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(RouteMap...)
	if err != nil {
		panic(err)
	}

	api.SetApp(router)
	log.Println("Start server at port: " + runPort)
	log.Fatal(http.ListenAndServe(":" +runPort, api.MakeHandler()))
}

func getWxConfig(w rest.ResponseWriter, r *rest.Request) {
	
}