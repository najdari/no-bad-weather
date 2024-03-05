run:
	go run cmd/cli/main.go 2>&1

test:
	go test -v ./...
