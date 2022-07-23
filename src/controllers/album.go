package controllers

import (
	"net/http"

	"github.com/WinIT23/web-service-gin/db"
	"github.com/WinIT23/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func DeleteAlbum(ctx *gin.Context) {
	d := db.GetDatabase()
	if id, ok := ctx.Params.Get("id"); ok {
		for i, a := range d.Albums {
			if a.ID == id {
				d.Albums = append(d.Albums[:i], d.Albums[i+1:]...)
				ctx.IndentedJSON(http.StatusOK, d.Albums)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, models.GenerateError(http.StatusNotFound, "Item not found."))
	} else {
		ctx.IndentedJSON(http.StatusInternalServerError, models.GenerateError(http.StatusInternalServerError, "Invalid request parameter."))
		return
	}
}

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, db.GetDatabase().Albums)
}

func GetAlbum(ctx *gin.Context) {
	d := db.GetDatabase()
	if id, ok := ctx.Params.Get("id"); ok {
		for _, a := range d.Albums {
			if a.ID == id {
				ctx.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, models.GenerateError(http.StatusNotFound, "Item not found."))
	} else {
		ctx.IndentedJSON(http.StatusInternalServerError, models.GenerateError(http.StatusInternalServerError, "Invalid request parameter."))
		return
	}
}

func PostAlbum(ctx *gin.Context) {
	var newAlbum models.Album
	d := db.GetDatabase()
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, models.GenerateError(http.StatusBadRequest, "Invalid JSON body."))
		return
	}

	d.Albums = append(d.Albums, newAlbum)

	ctx.IndentedJSON(http.StatusOK, d.Albums)
}

func UpdateAlbum(ctx *gin.Context) {
	var updatedAlbum models.Album
	d := db.GetDatabase()

	if err := ctx.BindJSON(&updatedAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, models.GenerateError(http.StatusBadRequest, "Invalid JSON body."))
		return
	}

	if id, ok := ctx.Params.Get("id"); ok {
		for i, a := range d.Albums {
			if a.ID == id {
				d.Albums[i] = updatedAlbum
				ctx.IndentedJSON(http.StatusOK, d.Albums)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, models.GenerateError(http.StatusNotFound, "Item not found."))
	} else {
		ctx.IndentedJSON(http.StatusInternalServerError, models.GenerateError(http.StatusInternalServerError, "Invalid request parameter."))
		return
	}
}
