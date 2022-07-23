package main

import (
	"github.com/WinIT23/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

const (
	Addr = "0.0.0.0:8080"
)

func main() {
	app := gin.Default()

	app.GET("/albums", handler.GetAlbums)
	app.GET("/albums/:id", handler.GetAlbum)
	app.POST("/albums", handler.PostAlbum)

	app.Run(Addr)
}
