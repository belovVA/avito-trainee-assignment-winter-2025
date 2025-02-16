build-up:
	docker-compose up

build-down:
	docker-compose down -v

test:
	 GO_ENV=test go test -coverprofile=coverage.out ./...

cover:
	 go tool cover -func=coverage.out

check:
	golanci-lint run
