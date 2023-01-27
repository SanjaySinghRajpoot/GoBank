package gapi

import (
	db "github.com/SanjaySinghRajpoot/GoBank/db/sqlc"
	"github.com/SanjaySinghRajpoot/GoBank/pb"
	"github.com/SanjaySinghRajpoot/GoBank/token"
	"github.com/gin-gonic/gin"
)

// GRPC service API
type Server struct {
	pb.UnimplementedGoBankServer
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(store *db.Store) *Server {
	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server
}
