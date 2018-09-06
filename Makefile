all: build generate

PACKAGE=github.com/openshift/cluster-dns-operator
MAIN_PACKAGE=$(PACKAGE)/cmd/cluster-dns-operator

BIN=$(lastword $(subst /, ,$(MAIN_PACKAGE)))
BINDATA=pkg/manifests/bindata.go

GOBINDATA_BIN=$(GOPATH)/bin/go-bindata

ENVVAR=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GOOS=linux
GO_BUILD_RECIPE=GOOS=$(GOOS) go build -o $(BIN) $(MAIN_PACKAGE)

build:
	$(GO_BUILD_RECIPE)

# Using "-modtime 1" to make generate target deterministic. It sets all file time stamps to unix timestamp 1
generate: $(GOBINDATA_BIN)
	go-bindata -mode 420 -modtime 1 -pkg manifests -o $(BINDATA) assets/...

$(GOBINDATA_BIN):
	go get -u github.com/jteeuwen/go-bindata/...

test:
	go test ./...

clean:
	go clean
	rm -f $(BIN)

.PHONY: all build generate test clean
