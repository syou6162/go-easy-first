COVERAGE = coverage.out

all: build

.PHONY: deps
deps:
	go get github.com/mattn/goveralls
	go get -u github.com/jteeuwen/go-bindata/...

.PHONY: bindata
bindata:
	go-bindata -ignore='\.gitkeep' data

.PHONY: build
build:
	go build -v

.PHONY: fmt
fmt:
	gofmt -s -w $$(git ls-files | grep -e '\.go$$' | grep -v -e vendor)

.PHONY: test
test:
	go test -v ./...

.PHONY: cover
cover:
	go test -v -cover -race -coverprofile=${COVERAGE}

.PHONY: vet
vet:
	go tool vet --all *.go

.PHONY: test-all
test-all: vet test
