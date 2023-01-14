#!/bin/sh

set -e 

echo "run db migration"
/app/migrate -path /app/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

echo "start the app"
exec "$@"