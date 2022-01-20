package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils/errors" //nolint:goimports

	"github.com/gin-gonic/gin" //nolint:goimports
	"github.com/minio/minio-go/v7"
)

type CreatChurchPartnerRequest struct {
	Name  string `json:"name" binding:"required"`
	Since string `json:"since" binding:"required"`
}

type GetAllChurchPartnersResponse struct {
	Total    int64                  `json:"total"`
	Partners []models.ChurchPartner `json:"news"`
}

func CreateChurchPartner(ctx *gin.Context) {

	type req CreatChurchPartnerRequest
	var reqData CreatChurchPartnerRequest
	var uploadedInfo minio.UploadInfo

	file, m, err := ctx.Request.FormFile("image")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	postData := req{
		Name:  ctx.PostForm("name"),
		Since: ctx.PostForm("since"),
	}
	reqData = CreatChurchPartnerRequest(postData)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := services.MinioService.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload church partner image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	uploadedInfo = uploadFile

	value := models.ChurchPartner{
		Name:  reqData.Name,
		Image: uploadedInfo.Key,
		Since: reqData.Since,
	}
	partner, errr := services.ChurchPartnersService.CreateChurchPartner(value)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create partner request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, partner)

}

func UpdateChurchPartner(ctx *gin.Context) {
	type req CreatChurchPartnerRequest
	var reqData CreatChurchPartnerRequest
	var uploadedInfo minio.UploadInfo

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
	file, m, err := ctx.Request.FormFile("image")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	postData := req{
		Name:  ctx.PostForm("name"),
		Since: ctx.PostForm("since"),
	}
	reqData = CreatChurchPartnerRequest(postData)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := services.MinioService.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	log.Println(data)

	uploadedInfo = uploadFile

	partner := models.ChurchPartner{
		Name:  reqData.Name,
		Image: uploadedInfo.Key,
		Since: reqData.Since,
	}
	errr := services.ChurchPartnersService.UpdateChurchPartner(uint(i), partner)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing update partner request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "partner model updated",
	})

}

func GetAllChurchPartner(ctx *gin.Context) {
	cacheData, errr := services.CacheService.GetPartnersList(context.Background(), "churchPartnersCache")
	log.Println(errr)

	if errr == nil {
		ctx.JSON(http.StatusOK, cacheData)
		return
	}
	partners, count, err := services.ChurchPartnersService.GetAllChurchPartner()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	s := make([]interface{}, len(partners))
	for i, v := range partners {
		s[i] = v
	}

	e := services.CacheService.SetPartnersList("churchPartnersCache", context.Background(), s)
	log.Println(e)

	data := GetAllChurchPartnersResponse{
		Total:    count,
		Partners: partners,
	}
	ctx.JSON(http.StatusOK, data)

}

func DeleteChurchPartner(ctx *gin.Context) {
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
	errr := services.ChurchPartnersService.DeleteChurchPartner(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted church partner",
	})

}

func GetSingleChurchPartner(ctx *gin.Context) {
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

	newsData, bg := services.CacheService.GetNews(context.Background(), "single-news")
	if bg == nil {
		log.Println(newsData)
		ctx.JSON(http.StatusOK, newsData)
		return
	}

	news, errr := services.ChurchPartnersService.GetSingleChurchPartner(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	//_ = services.CacheService.SetNews(context.Background(), news)
	ctx.JSON(http.StatusOK, news)

}