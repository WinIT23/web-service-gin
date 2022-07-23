package handler

import (
	"net/http"

	"github.com/WinIT23/web-service-gin/db"
	"github.com/WinIT23/web-service-gin/structures"

	"github.com/gin-gonic/gin"
)

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, db.GetDatabase().Albums)
}

func GetAlbum(ctx *gin.Context) {
	d := db.GetDatabase()
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	for _, a := range d.Albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, structures.GenerateError(http.StatusNotFound, "Item not found"))
}

func PostAlbum(ctx *gin.Context) {
	var newAlbum structures.Album
	d := db.GetDatabase()
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	d.Albums = append(d.Albums, newAlbum)

	ctx.IndentedJSON(http.StatusOK, db.GetDatabase().Albums)
}
