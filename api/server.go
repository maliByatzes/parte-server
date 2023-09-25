package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maliByatzes/parte-server/config"
	db "github.com/maliByatzes/parte-server/db/sqlc"
	"github.com/maliByatzes/parte-server/token"
	"github.com/maliByatzes/parte-server/worker"
)

type Server struct {
	config          config.Config
	store           db.Store
	router          *gin.Engine
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config config.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.Token.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a token maker: %v", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	server.setUpRouters()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
