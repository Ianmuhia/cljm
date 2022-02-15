package controllers

import (
	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type CreatSermonRequest struct {
	Title    string `json:"title" binding:"required"`
	Url      string `json:"url"  binding:"required"`
	DatePub  string `json:"date_pub" binding:"required"`
	Duration string `json:"duration"  binding:"required"`
}

type UpdateSermonRequest struct {
	Title    string `json:"title" `
	Url      string `json:"url" `
	DatePub  string `json:"date_pub" `
	Duration string `json:"duration"  `
}

type GetAllSermonsResponse struct {
	Total   int64            `json:"total"`
	Sermons []*models.Sermon `json:"sermons"`
}

func (r *Repository) CreateSermon(ctx *gin.Context) {
	type req CreatSermonRequest
	var reqData CreatSermonRequest
	var uploadedInfo minio.UploadInfo

	file, m, err := ctx.Request.FormFile("cover_image")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	postData := req{
		Title:    ctx.PostForm("title"),
		Url:      ctx.PostForm("url"),
		DatePub:  ctx.PostForm("date_pub"),
		Duration: ctx.PostForm("duration"),
	}
	reqData = CreatSermonRequest(postData)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	uploadedInfo = uploadFile
	sermonData := models.Sermon{
		Title:      reqData.Title,
		Url:        reqData.Url,
		DatePub:    reqData.DatePub,
		Duration:   reqData.Duration,
		CoverImage: uploadedInfo.Key,
	}
	sermon, err := r.sermonService.CreateSermon(sermonData)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing creating sermon post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, sermon)

}

func (r *Repository) UpdateSermon(ctx *gin.Context) {
	type req UpdateSermonRequest
	var reqData UpdateSermonRequest
	var uploadedInfo minio.UploadInfo

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
	file, m, err := ctx.Request.FormFile("cover_image")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	postData := req{
		Title:    ctx.PostForm("title"),
		Url:      ctx.PostForm("url"),
		DatePub:  ctx.PostForm("date_pub"),
		Duration: ctx.PostForm("duration"),
	}
	reqData = UpdateSermonRequest(postData)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	uploadedInfo = uploadFile
	sermonData := models.Sermon{
		Title:      reqData.Title,
		Url:        reqData.Url,
		DatePub:    reqData.DatePub,
		Duration:   reqData.Duration,
		CoverImage: uploadedInfo.Key,
	}

	errr := r.sermonService.UpdateSermon(uint(i), sermonData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing update partner request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sermon has been updated",
	})

}

func (r *Repository) GetAllSermons(ctx *gin.Context) {
	//cacheData, errr := services.CacheService.GetPartnersList(context.Background(), "churchPartnersCache")
	//log.Println(errr)
	//
	//if errr == nil {
	//	ctx.JSON(http.StatusOK, cacheData)
	//	return
	//}
	sermons, count, err := r.sermonService.GetAllSermon()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	//s := make([]interface{}, len(sermon))
	//for i, v := range sermon {
	//	s[i] = v
	//}
	//
	//e := services.CacheService.SetPartnersList("churchPartnersCache", context.Background(), s)
	//log.Println(e)

	data := GetAllSermonsResponse{
		Total:   count,
		Sermons: sermons,
	}
	ctx.JSON(http.StatusOK, data)

}

func (r *Repository) DeleteSermon(ctx *gin.Context) {
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
	errr := r.sermonService.DeleteSermon(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted sermon",
	})

}

func (r *Repository) GetSingleSermon(ctx *gin.Context) {
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

	sermon, errr := r.sermonService.GetSingleSermon(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	//_ = services.CacheService.SetNews(context.Background(), sermon)
	ctx.JSON(http.StatusOK, sermon)

}
