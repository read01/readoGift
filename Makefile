postgres:
	docker run --name pg15  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15.6

createdb:
	docker exec -it pg15 createdb --username=root --owner=root simple_bank


dropdb:
	docker exec -it pg15 dropdb simple_bank

migrateup:
	migrate -path ./db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	../bin/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc