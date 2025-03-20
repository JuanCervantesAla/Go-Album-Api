package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.AddAlbum)

	return router
}
