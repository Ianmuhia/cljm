package controllers

import (
	"context"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"

	"maranatha_web/models"
	"maranatha_web/services"
	"maranatha_web/utils/errors" //nolint:goimports

	"github.com/gin-gonic/gin" //nolint:goimports
	"github.com/minio/minio-go/v7"
)

type CreatEventsPostRequest struct {
	Title       string `json:"title" binding:"required"`
	SubTitle    string `json:"sub_title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	ScheduledOn string `json:"scheduledOn" binding:"required"`
}

type GetAllEventsResponse struct {
	Total  int64           `json:"total"`
	Events []*models.Event `json:"events"`
}

//TODO: Fix volunteer to event
func CreatEventsPost(ctx *gin.Context) {

	type req CreatEventsPostRequest
	var reqData CreatEventsPostRequest
	var uploadedInfo minio.UploadInfo

	//data := GetPayloadFromContext(ctx)
	file, m, err := ctx.Request.FormFile("cover_image")
	form, _ := ctx.MultipartForm()
	files := form.File["other_images"]

	for _, file := range files {
		log.Println(file.Filename)
	}

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	postData := req{
		Title:       ctx.PostForm("title"),
		SubTitle:    ctx.PostForm("sub_title"),
		Content:     ctx.PostForm("content"),
		ScheduledOn: ctx.PostForm("scheduledOn"),
	}
	reqData = CreatEventsPostRequest(postData)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := services.MinioService.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	uploadedInfo = uploadFile
	//TODO: Rework this.
	//user, err := services.UsersService.GetUserByEmail(data.Username)
	if err != nil {

		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	value := models.Event{
		Model:       gorm.Model{},
		OrganizerID: 0,
		//AuthorID:   user.ID,
		CoverImage:  uploadedInfo.Key,
		Title:       reqData.Title,
		SubTitle:    reqData.SubTitle,
		Content:     reqData.Content,
		ScheduledOn: "",
		//Jobs:        nil,/./
	}
	events, errr := services.EventsService.CreateEventsPost(value)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create events post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, events)

}

func UpdateEventsPost(ctx *gin.Context) {
	type req CreatEventsPostRequest
	var reqData CreatEventsPostRequest
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
	file, m, err := ctx.Request.FormFile("cover_image")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	//TODO:Create separate method to handle image upload
	postData := req{
		Title:    ctx.PostForm("title"),
		SubTitle: ctx.PostForm("sub_title"),
		Content:  ctx.PostForm("content"),
	}
	reqData = CreatEventsPostRequest(postData)
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
	eventsData := models.Event{
		CoverImage: uploadedInfo.Key,
		Title:      reqData.Title,
		SubTitle:   reqData.SubTitle,
		Content:    reqData.Content,
	}
	errr := services.EventsService.UpdateEventsPost(uint(i), eventsData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create events post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "events model updated",
	})

}

func GetAllEventsPost(ctx *gin.Context) {
	cacheData, errr := services.CacheService.GetEventsList(context.Background(), "events-list")

	if errr == nil {
		ctx.JSON(http.StatusOK, cacheData)
		return
	}
	events, count, err := services.EventsService.GetAllEventsPost()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	// _ = services.CacheService.SetEventsList(context.Background(), events)

	data := GetAllEventsResponse{
		Total:  count,
		Events: events,
	}
	ctx.JSON(http.StatusOK, data)

}

func GetAllEventsPostByAuthor(ctx *gin.Context) {
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
	events, count, errr := services.EventsService.GetAllEventsPostByAuthor(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	type GetAllEventsResponse2 struct {
		Total  int64           `json:"total"`
		Events []*models.Event `json:"events"`
	}
	data := GetAllEventsResponse2{
		Total:  count,
		Events: events,
	}
	ctx.JSON(http.StatusOK, data)

}

func DeleteEventsPost(ctx *gin.Context) {
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
	errr := services.EventsService.DeleteEventsPost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted events",
	})

}

func GetSingleEventsPost(ctx *gin.Context) {
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

	eventsData, bg := services.CacheService.GetEvents(context.Background(), "single-events")
	if bg == nil {
		log.Println(eventsData)
		ctx.JSON(http.StatusOK, eventsData)
		return
	}

	events, errr := services.EventsService.GetSingleEventsPost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	_ = services.CacheService.SetEvents(context.Background(), events)
	ctx.JSON(http.StatusOK, events)

}
