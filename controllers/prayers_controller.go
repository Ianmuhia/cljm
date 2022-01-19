package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils/errors"
	"net/http"
	"strconv"
)

type CreatePrayerPostRequest struct {
	Content string `json:"content" binding:"required"`
}

type GetAllPrayerResponse struct {
	Total  int64            `json:"total"`
	Prayer []*models.Prayer `json:"prayer"`
}

func CreatPrayerPost(ctx *gin.Context) {
	type req CreatePrayerPostRequest
	var reqData CreatePrayerPostRequest

	data := GetPayloadFromContext(ctx)
	postData := req{
		Content: ctx.PostForm("content"),
	}
	reqData = CreatePrayerPostRequest(postData)

	user, err := services.UsersService.GetUserByEmail(data.Username)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	value := models.Prayer{
		AuthorID: user.ID,
		Content:  reqData.Content,
	}
	prayer, errr := services.PrayerService.CreatePrayerPost(value)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create prayer post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, prayer)

}

func UpdatePrayerPost(ctx *gin.Context) {
	type req CreatePrayerPostRequest
	var reqData CreatePrayerPostRequest

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
	reqData = CreatePrayerPostRequest(postData)

	log.Println(data)
	prayerData := models.Prayer{
		Content: reqData.Content,
	}
	errr := services.PrayerService.UpdatePrayerPost(uint(i), prayerData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create prayer post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "prayer model updated",
	})

}

func GetAllPrayerPost(ctx *gin.Context) {
	cacheData, errr := services.CacheService.GetPrayerList(context.Background(), "prayers-list")

	if errr == nil {
		ctx.JSON(http.StatusOK, cacheData)
		return
	}
	prayer, count, err := services.PrayerService.GetAllPrayerPost()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	data := GetAllPrayerResponse{
		Total:  count,
		Prayer: prayer,
	}
	ctx.JSON(http.StatusOK, data)

}

func GetAllPrayerPostByAuthor(ctx *gin.Context) {
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
	prayer, count, errr := services.PrayerService.GetAllPrayerPostByAuthor(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	type GetAllPrayerResponse2 struct {
		Total  int64            `json:"total"`
		Prayer []*models.Prayer `json:"prayer"`
	}
	data := GetAllPrayerResponse2{
		Total:  count,
		Prayer: prayer,
	}
	ctx.JSON(http.StatusOK, data)

}

func DeletePrayerPost(ctx *gin.Context) {
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
	errr := services.PrayerService.DeletePrayerPost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted prayer",
	})

}

func GetSinglePrayerPost(ctx *gin.Context) {
	//TODO:Create method for getting and converting this id
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

	prayerData, bg := services.CacheService.GetPrayer(context.Background(), "single-prayer")
	if bg == nil {
		log.Println(prayerData)
		ctx.JSON(http.StatusOK, prayerData)
		return
	}

	prayer, errr := services.PrayerService.GetSinglePrayerPost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	_ = services.CacheService.SetPrayer(context.Background(), prayer)
	ctx.JSON(http.StatusOK, prayer)

}
