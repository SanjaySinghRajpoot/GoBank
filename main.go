package main

import (
	"database/sql"
	"log"

	"github.com/SanjaySinghRajpoot/GoBank/api"
	db "github.com/SanjaySinghRajpoot/GoBank/db/sqlc"
	"github.com/SanjaySinghRajpoot/GoBank/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	dbDri := string(config.DBDriver)
	dbSou := string(config.DBSource)

	conn, err := sql.Open(dbDri, dbSou)
	if err != nil {
		log.Fatal("cannot connect", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
