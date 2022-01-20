package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils/errors"
)

type CreateGenrePostRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetAllGenreResponse struct {
	Total int64           `json:"total"`
	Genre []*models.Genre `json:"genre"`
}

func CreatGenrePost(ctx *gin.Context) {
	var req CreateGenrePostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	postData := models.Genre{
		Name: req.Name,
	}

	genre, errr := services.GenreService.CreateGenrePost(postData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create genre post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, genre)

}

func UpdateGenrePost(ctx *gin.Context) {

	id := ctx.Query("id")
	value, _ := strconv.ParseInt(id, 10, 32)
	if id == "" || value == 0 {
		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		data := errors.NewBadRequestError("Provide an id to the request.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return

	}

	var req CreateGenrePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	genreData := models.Genre{
		Name: req.Name,
	}
	errr := services.GenreService.UpdateGenrePost(uint(i), genreData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create genre post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "genre model updated",
	})

}

func GetAllGenrePost(ctx *gin.Context) {
	cacheData, errr := services.CacheService.GetGenreList(context.Background(), "genres-list")
	if errr == nil {
		ctx.JSON(http.StatusOK, cacheData)
		return
	}
	genre, count, err := services.GenreService.GetAllGenrePost()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	data := GetAllGenreResponse{
		Total: count,
		Genre: genre,
	}
	ctx.JSON(http.StatusOK, data)

}

func DeleteGenrePost(ctx *gin.Context) {
	id := ctx.Query("id")
	value, _ := strconv.ParseInt(id, 10, 32)
	if id == "" || value == 0 {
		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		data := errors.NewBadRequestError("Provide an id to the request.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return

	}
	errr := services.GenreService.DeleteGenrePost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted genre",
	})

}

func GetSingleGenrePost(ctx *gin.Context) {
	id := ctx.Query("id")
	value, _ := strconv.ParseInt(id, 10, 32)
	if id == "" || value == 0 {
		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		data := errors.NewBadRequestError("Provide an id to the request.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return

	}

	genreData, bg := services.CacheService.GetGenre(context.Background(), "single-genre")
	if bg == nil {
		log.Println(genreData)
		ctx.JSON(http.StatusOK, genreData)
		return
	}

	genre, errr := services.GenreService.GetSingleGenrePost(strconv.FormatUint(i, 10)) //nolint:govet
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	_ = services.CacheService.SetGenre(context.Background(), genre)
	ctx.JSON(http.StatusOK, genre)

}
