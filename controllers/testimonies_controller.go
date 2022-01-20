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

type CreateTestimonyRequest struct {
	Content string `json:"content" binding:"required"`
}

type GetAllTestimoniesResponse struct {
	Total       int64                 `json:"total"`
	Testimonies []*models.Testimonies `json:"testimonies"`
}

func CreateTestimony(ctx *gin.Context) {
	type req CreateTestimonyRequest
	var reqData CreateTestimonyRequest

	data := GetPayloadFromContext(ctx)
	postData := req{
		Content: ctx.PostForm("content"),
	}
	reqData = CreateTestimonyRequest(postData)

	user, err := services.UsersService.GetUserByEmail(data.Username)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	value := models.Testimonies{
		AuthorID: user.ID,
		Content:  reqData.Content,
	}
	prayer, errr := services.TestimoniesService.CreateTestimony(value)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create testimonies post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, prayer)

}

func UpdateTestimony(ctx *gin.Context) {
	type req CreateTestimonyRequest
	var reqData CreateTestimonyRequest

	data := GetPayloadFromContext(ctx)
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

	//TODO:Create separate method to handle image upload
	postData := req{

		Content: ctx.PostForm("content"),
	}
	reqData = CreateTestimonyRequest(postData)

	log.Println(data)
	testimoniesData := models.Testimonies{

		Content: reqData.Content,
	}
	errr := services.TestimoniesService.UpdateTestimony(uint(i), testimoniesData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create testimonies post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "testimonies model updated",
	})

}

func GetAllTestimonies(ctx *gin.Context) {
	cacheData, errr := services.CacheService.GetTestimoniesList(context.Background(), "testimonies-list")

	if errr == nil {
		ctx.JSON(http.StatusOK, cacheData)
		return
	}
	testimonies, count, err := services.TestimoniesService.GetAllTestimonies()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	data := GetAllTestimoniesResponse{
		Total:       count,
		Testimonies: testimonies,
	}
	ctx.JSON(http.StatusOK, data)

}

func GetAllTestimoniesByAuthor(ctx *gin.Context) {
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
	testimonies, count, errr := services.TestimoniesService.GetAllTestimoniesByAuthor(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	type GetAllTestimoniesResponse2 struct {
		Total       int64                 `json:"total"`
		Testimonies []*models.Testimonies `json:"testimonies"`
	}
	data := GetAllTestimoniesResponse2{
		Total:       count,
		Testimonies: testimonies,
	}
	ctx.JSON(http.StatusOK, data)

}

func DeleteTestimony(ctx *gin.Context) {
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
	errr := services.TestimoniesService.DeleteTestimony(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted testimonies",
	})

}

func GetSingleTestimony(ctx *gin.Context) {

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

	testimoniesData, bg := services.CacheService.GetTestimonies(context.Background(), "single-testimonies")
	if bg == nil {
		log.Println(testimoniesData)
		ctx.JSON(http.StatusOK, testimoniesData)
		return
	}

	testimonies, errr := services.TestimoniesService.GetSingleTestimony(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	_ = services.CacheService.SetTestimonies(context.Background(), testimonies)
	ctx.JSON(http.StatusOK, testimonies)

}
