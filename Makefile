postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root POSTGRES_PASSWORD=secret -d postgres:12-alpine

createddb:
    docker exec -it postgres createddb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres dropdb simple_bank

migrateup:
    migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateRDS:
    migrate -path db/migrate -database "postgres://root:Dyy6hA8TsCP5uPudwEsG@go-bank.cxfbvzltflk5.ap-south-1.rds.amazonaws.com:5432/simple_bank" -verbose up

migratedown:
        migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
        migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
 
migratedown1:
        migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
 
db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
   go run main.go

db_migration:
   migrate create -ext sql db/migration -seq <migration_name>

proto:
   rm -f pb/*.go
   protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

proto(grpc+http):
   rm -f pb/*.go
   rm -f doc/swagger/*.swagger.json
   protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=go_bank \
    proto/*.proto

evans: 
  evans --host localhost --port  9090 -r repl

.PHONY: network postgres server