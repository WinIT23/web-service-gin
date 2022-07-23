package main

import (
	"fmt"
	"os"

	"github.com/WinIT23/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/albums", handler.GetAlbums)
	app.GET("/albums/:id", handler.GetAlbum)
	app.POST("/albums", handler.PostAlbum)

	fmt.Println(getHost())
	app.Run(getHost())
}

func getHost() string {
	host := os.Getenv("GO_RUN_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("GO_RUN_PORT")
	if port == "" {
		port = "8080"
	}
	return host + ":" + port
}
