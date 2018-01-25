package main

import "os"

var (
	wxAppId = os.Getenv("WXAppId")
	wxAppSecret = os.Getenv("WXAppSecret")

	runPort = os.Getenv("RunPort")
)
