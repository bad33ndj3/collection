# Format Go code using gofumpt
fmt:
	GO111MODULE=off go get mvdan.cc/gofumpt
	gofumpt -s -w .

# Run tests and generate coverage report
test: reports/coverage.out
	go test -v -coverprofile=reports/coverage.out ./...

# Open coverage report in browser
coverage:
	go tool cover -html=reports/coverage.out

# Run golangci-lint to check for issues
lint:
	CGO_ENABLED=0 golangci-lint run

# Create reports directory if it doesn't exist
reports/coverage.out:
	mkdir reports
	touch reports/coverage.out
