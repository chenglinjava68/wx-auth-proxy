export GOPATH:=$(shell pwd)

dep:
	go get -d -v ./...

fmt:
	go fmt ./...

install: dep
	go install main

clean:
	go clean -i -r ./...