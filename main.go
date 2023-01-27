package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/SanjaySinghRajpoot/GoBank/api"
	db "github.com/SanjaySinghRajpoot/GoBank/db/sqlc"
	"github.com/SanjaySinghRajpoot/GoBank/gapi"
	"github.com/SanjaySinghRajpoot/GoBank/pb"
	"github.com/SanjaySinghRajpoot/GoBank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// runGinServer(config, *store)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store *db.Store) {

	server := gapi.NewServer(store)
	// if err != nil {
	// 	log.Fatal("cannot start server: ", err)
	// }

	grpcServer := grpc.NewServer()
	pb.RegisterGoBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server ")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGinServer(config util.Config, store *db.Store) {
	server := api.NewServer(config, *store)

	err := server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
