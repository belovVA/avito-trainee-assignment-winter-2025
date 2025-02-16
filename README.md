# Avito-shop
 avito-trainee-assignment-winter-2025

## Запуск сервиса
Необходимо запустить следующие команды
```
1. make check
2. make test
3. make build-up
```
Для завершения работы сервиса введите: \
`make build-down`

Структура проекта:
```
.
├── 1.txt
├── cmd
│   └── avito-coin-service
│       └── main.go
├── config
│   ├── jwt.go
│   └── postgre.go
├── coverage.out
├── docker-compose.yaml
├── Dockerfile
├── golangci.yaml
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── app.go
│   ├── database
│   │   └── postgres.go
│   ├── handler
│   │   ├── info.go
│   │   ├── purchase.go
│   │   ├── transaction.go
│   │   └── user.go
│   ├── middleware
│   │   ├── auth.go
│   │   └── jwt.go
│   ├── model
│   │   ├── info.go
│   │   ├── merch.go
│   │   ├── purchase.go
│   │   ├── transaction.go
│   │   └── user.go
│   ├── repository
│   │   ├── merch.go
│   │   ├── purchase.go
│   │   ├── transaction.go
│   │   └── user.go
│   └── service
│       ├── info.go
│       ├── purchase.go
│       ├── tests
│       │   ├── info_test.go
│       │   ├── purchase_test.go
│       │   ├── transaction_test.go
│       │   └── user_test.go
│       ├── transaction.go
│       └── user.go
├── Makefile
├── migrations
│   ├── down
│   │   ├── 00001_users_table.down.sql
│   │   ├── 00002_transaction_table.down.sql
│   │   ├── 00003_merch_table.down.sql
│   │   └── 00004_purchases_table.down.sql
│   └── up
│       ├── 00001_users_table.up.sql
│       ├── 00002_transaction_table.up.sql
│       ├── 00003_merch_table.up.sql
│       └── 00004_purchases_table.up.sql
├── mocks
│   ├── merch.go
│   ├── purchase.go
│   ├── transaction.go
│   └── user.go
└── README.md

17 directories, 49 files
```
