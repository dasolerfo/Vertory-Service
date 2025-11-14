package api

import (
	"fmt"

	//db "simplebank/db/model"
	//"simplebank/factory"
	factory "github.com/dasolerfo/Vertory-Service/factory"
	token "github.com/dasolerfo/Vertory-Service/token"

	db "github.com/dasolerfo/Vertory-Service/db/model"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     factory.Config
	tokenMaker token.Maker
	store      db.Querier
	router     *gin.Engine
}

func NewServer(config factory.Config, store db.Querier) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}
	server := &Server{store: store,
		tokenMaker: tokenMaker,
		config:     config}

	/*if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}*/

	server.Router()
	return server, nil
}

func (server *Server) Router() {

	router := gin.Default()
	server.router = router
	/*
	   router.POST("/owner", server.createOwner)
	   router.POST("/owner/login", server.loginOwner)

	   authRoutes := router.Group("/").Use(AuthMiddleware(server.tokenMaker))

	   authRoutes.POST("/token/renew", server.renewAccessToken)

	   authRoutes.POST("/accounts", server.createAccount)
	   authRoutes.GET("/accounts/:id", server.getAccount)
	   authRoutes.GET("/accounts", server.ListAccount)

	   authRoutes.POST("/transfers", server.createTransferMoneyAPI)

	   server.router = router
	*/
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
