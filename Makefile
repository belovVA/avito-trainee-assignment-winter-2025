run:
	go run cmd/avito-shop/main.go

test:
	GO_ENV=test go test ./internal/service/tests/...

cover:
	go test ./... -v -coverpkg=./...