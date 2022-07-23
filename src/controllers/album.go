package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/WinIT23/web-service-gin/configs"
	"github.com/WinIT23/web-service-gin/models"
	"github.com/WinIT23/web-service-gin/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "albums"

var collection *mongo.Collection = configs.GetCollection(configs.GetMongoClient(), collectionName)
var validate = validator.New()

func DeleteAlbum(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	albumId, ok := ctx.Params.Get("id")
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, responses.GenerateResponse(http.StatusBadRequest, "error", map[string]interface{}{"data": "Invalid ID provided"}))
		return
	}

	id, _ := primitive.ObjectIDFromHex(albumId)
	result, err := collection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, responses.GenerateResponse(http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()}))
		return
	}

	if result.DeletedCount == 0 {
		ctx.IndentedJSON(http.StatusNotFound, responses.GenerateResponse(http.StatusNotFound, "error", map[string]interface{}{"data": "Document not found"}))
		return
	}
	ctx.IndentedJSON(http.StatusCreated, responses.GenerateResponse(http.StatusCreated, "success", map[string]interface{}{"data": result}))
}

func GetAlbums(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var albums []models.Album
	defer cancel()

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, responses.GenerateResponse(http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()}))
		return
	}

	defer cursor.Close(c)
	for cursor.Next(c) {
		var album models.Album
		if err := cursor.Decode(&album); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, responses.GenerateResponse(http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()}))
			return
		}
		albums = append(albums, album)
	}
	ctx.IndentedJSON(http.StatusOK, responses.GenerateResponse(http.StatusOK, "success", map[string]interface{}{"data": albums}))
}

func GetAlbum(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var album models.Album
	defer cancel()

	albumId, ok := ctx.Params.Get("id")
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, responses.GenerateResponse(http.StatusBadRequest, "error", map[string]interface{}{"data": "Invalid ID provided"}))
		return
	}

	id, _ := primitive.ObjectIDFromHex(albumId)

	if err := collection.FindOne(c, bson.M{"_id": id}).Decode(&album); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, responses.GenerateResponse(http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()}))
		return
	}

	ctx.IndentedJSON(http.StatusOK, responses.GenerateResponse(http.StatusOK, "success", map[string]interface{}{"data": album}))
}

func CreateAlbum(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var album models.Album
	defer cancel()

	if err := ctx.BindJSON(&album); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, responses.GenerateResponse(http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()}))
		return
	}

	if err := validate.Struct(&album); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, responses.GenerateResponse(http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()}))
		return
	}

	newAlbum := models.Album{
		ID:     primitive.NewObjectID(),
		Artist: album.Artist,
		Title:  album.Title,
		Price:  album.Price,
	}

	fmt.Printf("%#v", newAlbum)

	result, err := collection.InsertOne(c, newAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, responses.GenerateResponse(http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()}))
		return
	}

	ctx.IndentedJSON(http.StatusCreated, responses.GenerateResponse(http.StatusCreated, "success", map[string]interface{}{"data": result}))
}

func UpdateAlbum(ctx *gin.Context) {}
