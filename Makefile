binaryName:="grt"
buildRevision:= `git rev-parse --short HEAD`

all: build
build:
	go build -ldflags="-X main.buildRevision=$(buildRevision) -s -w" -o $(binaryName) main.go; upx $(binaryName)

.PHONY: test
test:
	go test -race -v

.PHONY: clean
clean:
	rm -f $(binaryName)

.PHONY: style-fix
style-fix:
	gofmt -w .

.PHONE: lint
lint:
	golangci-lint run

.PHONY: upgrade
upgrade:
	go mod tidy
	go get -u all ./...