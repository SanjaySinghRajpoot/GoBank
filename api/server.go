package api

import (
	db "github.com/SanjaySinghRajpoot/GoBank/db/sqlc"
	"github.com/SanjaySinghRajpoot/GoBank/token"
	"github.com/SanjaySinghRajpoot/GoBank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewServer(config util.Config, store db.Store) *Server {
	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("curreny", validCurrency)
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/token/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker)) // adding auth Routes for verification

	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts", server.listAccount)

	authRoutes.POST("/tranfers", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
