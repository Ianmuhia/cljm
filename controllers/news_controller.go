package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //nolint:goimports
	"github.com/minio/minio-go/v7"
	"maranatha_web/controllers/token"
	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils/errors"
)

type CreatNewsPostRequest struct {
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"sub_title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type GetAllNewsResponse struct {
	Total int64         `json:"total"`
	News  []models.News `json:"news"`
}

func CreatNewsPost(ctx *gin.Context) {
	type req CreatNewsPostRequest
	var req_data CreatNewsPostRequest

	var uploadedInfo minio.UploadInfo
	file, m, err := ctx.Request.FormFile("cover_image")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	fmt.Println(file, m.Header, m.Filename, m.Size)
	post_data := req{
		Title:    ctx.PostForm("title"),
		SubTitle: ctx.PostForm("sub_title"),
		Content:  ctx.PostForm("content"),
	}
	req_data = CreatNewsPostRequest(post_data)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := services.MinioService.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	log.Println(uploadFile)
	uploadedInfo = uploadFile

	payload, exists := ctx.Get("authorization_payload")
	if !exists {
		restErr := errors.NewBadRequestError("could not get auth_payload from context")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	data := payload.(*token.Payload)
	user, err := services.UsersService.GetUserByEmail(data.Username)
	if err != nil {
		//log.Println(user)
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	//TODO:upload news images
	//TODO:Work on user profile, reset password , forgot password
	value := models.News{
		AuthorID:   user.ID,
		CoverImage: uploadedInfo.Key,
		Title:      req_data.Title,
		SubTitle:   req_data.SubTitle,
		Content:    req_data.Content,
	}
	news, errr := services.NewsService.CreateNewsPost(value)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create news post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, news)

}

func GetAllNewsPost(ctx *gin.Context) {
	news, count, err := services.NewsService.GetAllNewsPost()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	data := GetAllNewsResponse{
		Total: count,
		News:  news,
	}
	ctx.JSON(http.StatusOK, data)

}

func DeleteNewsPost(ctx *gin.Context) {
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
	errr := services.NewsService.DeleteNewsPost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted news",
	})

}

func GeSingleNewsPost(ctx *gin.Context) {
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
	news, errr := services.NewsService.GetSingleNewsPost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, news)

}
