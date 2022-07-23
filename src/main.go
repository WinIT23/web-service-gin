package main

import (
	"os"

	"github.com/WinIT23/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	Addr := os.Getenv("GO_RUN_HOST") + ":" + os.Getenv("GO_RUN_PORT")
	app := gin.Default()

	app.GET("/albums", handler.GetAlbums)
	app.GET("/albums/:id", handler.GetAlbum)
	app.POST("/albums", handler.PostAlbum)

	app.Run(Addr)
}
