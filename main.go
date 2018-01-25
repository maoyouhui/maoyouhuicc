package main

import "github.com/ant0ine/go-json-rest/rest"
import (
	"net/http"
	"log"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/jssdk"
	"gopkg.in/chanxuehong/wechat.v2/util"
	"time"
	"strconv"
)

var accessTokenServer core.AccessTokenServer

func main()  {

	accessTokenServer = core.NewDefaultAccessTokenServer(wxAppId, wxAppSecret, nil)
	accessTokenServer.RefreshToken("")

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

	//Access token
	accessTokenServer.RefreshToken("")

	//js ticket
	var ticketServer = jssdk.NewDefaultTicketServer(core.NewClient(accessTokenServer, nil))

	jsapiTicket, err := ticketServer.Ticket()

	if err != nil {
		panic(err)
	}

	nonceStr := util.NonceStr()
	timestamp := time.Now().Unix()
	url := r.URL.Query().Get("url")

	signature := jssdk.WXConfigSign(jsapiTicket, nonceStr, strconv.Itoa(int(timestamp)), url)

	var ret = &map[string]interface{}{
		"nonceStr": nonceStr,
		"timestamp": timestamp,
		"signature": signature,
	}

	w.WriteJson(ret)
}