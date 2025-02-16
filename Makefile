run:
	go run cmd/avito-shop/main.go

test:
	GO_ENV=test go test -v -coverprofile=coverage.out ./...

cover:
	GO_ENV=test go tool cover -func=coverage.out
