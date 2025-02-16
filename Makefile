run:
	go run cmd/avito-shop/main.go

test:
	 go test -coverprofile=coverage.out ./...

cover:
	 go tool cover -func=coverage.out
