postgres:
	docker run --name pg15  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15.6

createdb:
	docker exec -it pg15 createdb --username=root --owner=root simple_bank


dropdb:
	docker exec -it pg15 dropdb simple_bank

migrateup:
	migrate -path ./db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path ./db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

pg:
	docker exec -it pg15 psql -U root -d simple_bank 

server:
	go run main.go 

mock:
	mockgen -package mockdb  -destination db/mock/store.go  readoGift/db/sqlc Store 

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test pg server mock migrateup1 migratedown1