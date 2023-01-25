BINARY_NAME=ynab-lazy-import
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	go build -o $(BINARY_NAME) -v

fmt:
	GO111MODULE=off go get mvdan.cc/gofumpt
	gofumpt -s -w .

test: reports/coverage.out
	go test -v -coverprofile=reports/coverage.out ./...

reports/coverage.out:
	mkdir reports
	touch reports/coverage.out

coverage:
	go tool cover -html=reports/coverage.out

lint:
	CGO_ENABLED=0 golangci-lint run
