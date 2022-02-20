package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
	"net/http"
	"strconv"
)

type CreatPodcastPostRequest struct {
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Cast        string `json:"cast"`
	Description string `json:"description"`
	PodcastUrl  string `json:"podcastUrl"`
}

type GetAllPodcastResponse struct {
	Total   int               `json:"total"`
	Podcast []*models.Podcast `json:"podcast"`
}

func (r *Repository) CreatePodcast(ctx *gin.Context) {
	type req CreatPodcastPostRequest
	var reqData CreatPodcastPostRequest
	var uploadedInfo minio.UploadInfo

	data := r.GetPayloadFromContext(ctx)
	file, m, err := ctx.Request.FormFile("cover_image")

	log.Println(file)
	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	postData := req{
		Title:       ctx.PostForm("title"),
		SubTitle:    ctx.PostForm("sub_title"),
		Cast:        ctx.PostForm("cast"),
		Description: ctx.PostForm("description"),
		PodcastUrl:  ctx.PostForm("podcastUrl"),
	}
	reqData = CreatPodcastPostRequest(postData)

	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	uploadedInfo = uploadFile
	value := models.Podcast{
		AuthorID:    data.ID,
		CoverImage:  uploadedInfo.Key,
		Title:       reqData.Title,
		SubTitle:    reqData.SubTitle,
		Cast:        reqData.Cast,
		Description: reqData.Description,
		PodcastUrl:  reqData.PodcastUrl,
	}
	podcast, err := r.podcastService.CreatePodcast(value)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing create podcast post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	resp := NewStatusCreatedResponse("Podcast created successfully", podcast)

	ctx.JSON(resp.Status, resp)

}

func (r *Repository) UpdatePodcast(ctx *gin.Context) {
	type req CreatPodcastPostRequest
	var reqData CreatPodcastPostRequest
	var uploadedInfo minio.UploadInfo

	data := r.GetPayloadFromContext(ctx)
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
		Title:       ctx.PostForm("title"),
		SubTitle:    ctx.PostForm("sub_title"),
		Cast:        ctx.PostForm("cast"),
		Description: ctx.PostForm("description"),
		PodcastUrl:  ctx.PostForm("podcastUrl")}
	reqData = CreatPodcastPostRequest(postData)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	log.Println(data)
	uploadedInfo = uploadFile
	podcastData := models.Podcast{
		CoverImage:  uploadedInfo.Key,
		Title:       reqData.Title,
		SubTitle:    reqData.SubTitle,
		Cast:        reqData.Cast,
		Description: reqData.Description,
		PodcastUrl:  reqData.PodcastUrl,
	}
	errr := r.podcastService.UpdatePodcast(uint(i), podcastData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create podcast post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "podcast model updated",
	})

}

func (r *Repository) GetAllPodcast(ctx *gin.Context) {

	data, count, err := r.podcastService.GetAllPodcast()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	// _ = services.CacheService.SetPodcastList(context.Background(), podcast)

	podcast := GetAllPodcastResponse{
		Total:   count,
		Podcast: data,
	}

	resp := NewStatusOkResponse("Got podcast successfully", podcast)

	ctx.JSON(resp.Status, resp)

}

func (r *Repository) GetAllPodcastByAuthor(ctx *gin.Context) {
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
	podcast, count, err := r.podcastService.GetAllPodcastByAuthor(uint(i))
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	data := GetAllPodcastResponse{
		Total:   count,
		Podcast: podcast,
	}
	resp := NewStatusOkResponse("All podcast by this author.", data)
	ctx.JSON(resp.Status, resp)
}

func (r *Repository) DeletePodcast(ctx *gin.Context) {
	id := ctx.Query("id")
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		data := errors.NewBadRequestError("Provide an id to the request.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return

	}
	err = r.podcastService.DeletePodcast(uint(i))
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	resp := NewDeleteResponse("Successfully deleted podcast", nil)
	ctx.JSON(resp.Status, resp)

}

func (r *Repository) GetSinglePodcast(ctx *gin.Context) {
	id := ctx.Query("id")

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		data := errors.NewBadRequestError("Provide an id to the request.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	podcast, err := r.podcastService.GetSinglePodcast(uint(i))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			data := errors.NewNotFoundError("Not found")
			ctx.JSON(data.Status, data)
			ctx.Abort()
			return
		}
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	resp := NewStatusOkResponse("Get single new success.", podcast)
	//_ = services.CacheService.SetPodcast(context.Background(), podcast)
	ctx.JSON(resp.Status, resp)

}
