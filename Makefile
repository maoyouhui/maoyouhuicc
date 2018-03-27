SHELL := /bin/bash

build :

	go build -o myhcc  *.go

run : build

	export RunPort=3000 && source .env && ./myhcc $(RUN_ARGS)

run-bbs : build

	export RunPort=3001 && source .bbs.env && ./myhcc &(RUN_ARGS)

stop :
	pkill myhcc

list-tasks:

	grep "\"Task.*\"" main.go | cut -d : -f 1
