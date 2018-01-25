SHELL := /bin/bash

build :

	go build -o myhcc  *.go

run :

	export RunPort=3000 && source .env && ./myhcc $(RUN_ARGS)

list-tasks:

	grep "\"Task.*\"" main.go | cut -d : -f 1