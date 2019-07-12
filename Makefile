all:
	@echo "please read Makefile source or README to see available commands"

tools:
	GO111MODULE=off go get -u github.com/webrpc/webrpc/cmd/webrpc-gen
	GO111MODULE=off go get -u github.com/goware/webify

generate: generate-server generate-client

generate-server:
	webrpc-gen -schema=proto.ridl -target=go -pkg=main -client -server -out=./proto.gen.go

generate-client:
	webrpc-gen -schema=proto.ridl -target=js -extra=noexports -client -out=./client-js/remotelog.js
	webrpc-gen -schema=proto.ridl -target=ts -client -out=./client-ts/remotelog.ts

run:
	go run .

install:
	go install .
