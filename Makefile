postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root POSTGRES_PASSWORD=secret -d postgres:12-alpine

createddb:
    docker exec -it postgres createddb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres dropdb simple_bank

migrateup:
    migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
   sqlc generate

.PHONY: postgres createddb dropdb migrateup