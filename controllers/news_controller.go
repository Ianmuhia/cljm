package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"maranatha_web/controllers/token"
	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils/errors"
)

type CreatNewsPostRequest struct {
	//CoverImage string `json:"cover_image" binding:"required"`
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"sub_title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type CreatNewsPostResponse struct {
	Message string `json:"message"`
}

func CreatNewsPost(ctx *gin.Context) {
	var req CreatNewsPostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
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
		return
	}
	value := models.News{
		AuthorID:   user.ID,
		CoverImage: "",
		Title:      req.Title,
		SubTitle:   req.SubTitle,
		Content:    req.Content,
	}
	news, _ := services.NewsService.CreateNewsPost(value)

	log.Println(news)
	ctx.JSON(http.StatusCreated, data)

}
