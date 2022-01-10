package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"maranatha_web/controllers/token"
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
	var payload token.Payload
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	payloadd, exists := ctx.Get("authorization_payload")
	if !exists {
		restErr := errors.NewBadRequestError("could not get auth_payload from context")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	dc := payloadd.(*token.Payload)
	log.Println(payload)
	log.Println(dc)
	log.Println(ctx.Get("authorization_payload"))

	ctx.JSON(http.StatusCreated, "payload")

}
