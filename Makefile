postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root POSTGRES_PASSWORD=secret -d postgres:12-alpine

createddb:
    docker exec -it postgres createddb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres dropdb simple_bank

.PHONY: 