all: build

.PHONY: deps
deps:
	go get github.com/mattn/goveralls

.PHONY: build
build:
	go build
