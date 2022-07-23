package routers

import (
	"github.com/WinIT23/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoute(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbum)
	router.POST("/albums", controllers.CreateAlbum)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
}
