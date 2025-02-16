build-up:
	docker-compose up -d

build-down:
	docker-compose down -v
	
test:
	 go test -coverprofile=coverage.out ./...

cover:
	 go tool cover -func=coverage.out
