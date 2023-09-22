package api

import "github.com/gin-gonic/gin"

func (server *Server) setUpRouters() {
	router := gin.Default()

	// Non protetced routes
	router.POST("/users/register", server.createUser)

	server.router = router
}
