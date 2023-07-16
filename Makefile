build-app:
	go build -o bin/server main.go

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

