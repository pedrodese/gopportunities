package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize router
	router := gin.Default()
	//Initialize router
	initializeRoutes(router)
	//Run the server
	router.Run()
}
