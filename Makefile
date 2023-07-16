build-app:
	go build -o bin/server main.go

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

createdb:
	docker exec -it postgres_container --username=root --owner=root movie

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/movie?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/movie?sslmode=disable" -verbose down

.PHONY: build-app migrate createdb migrateup migratedown