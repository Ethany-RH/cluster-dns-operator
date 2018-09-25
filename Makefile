all: generate build

PACKAGE=github.com/openshift/cluster-dns-operator
MAIN_PACKAGE=$(PACKAGE)/cmd/cluster-dns-operator

BIN=$(lastword $(subst /, ,$(MAIN_PACKAGE)))
BINDATA=pkg/manifests/bindata.go
TEST_BINDATA=test/manifests/bindata.go

vpath bin/go-bindata $(GOPATH)
GOBINDATA_BIN=bin/go-bindata

ENVVAR=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GOOS=linux
GO_BUILD_RECIPE=GOOS=$(GOOS) go build -o $(BIN) $(MAIN_PACKAGE)

build:
	$(GO_BUILD_RECIPE)

# Using "-modtime 1" to make generate target deterministic. It sets all file time stamps to unix timestamp 1
generate: $(GOBINDATA_BIN)
	go-bindata -mode 420 -modtime 1 -pkg manifests -o $(BINDATA) assets/...
	go-bindata -mode 420 -modtime 1 -pkg manifests -o $(TEST_BINDATA) test/assets/...

$(GOBINDATA_BIN):
	go get -u github.com/jteeuwen/go-bindata/...

test:
	go test ./...

test-integration:
	go test -v -tags integration ./test/integration

clean:
	go clean
	rm -f $(BIN)

.PHONY: all build generate test test-integration deps-integration clean
