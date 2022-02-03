package controllers

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"time"
//
//
//
//

// 	"github.com/gin-gonic/gin" //nolint:goimports
// 	"github.com/filestorage/filestorage-go/v7"

// 	"maranatha_web/models"
// 	"maranatha_web/services"
// 	"maranatha_web/utils/date_utils"
// 	"maranatha_web/utils/errors" //nolint:goimports
// )

// type CreatEventsPostRequest struct {
// 	Title       string `json:"title" binding:"required"`
// 	SubTitle    string `json:"sub_title" binding:"required"`
// 	Content     string `json:"content" binding:"required"`
// 	ScheduledOn string `json:"scheduledOn" binding:"required"`
// }

// type GetAllEventsResponse struct {
// 	Total  int64                 `json:"total"`
// 	Events []*models.ChurchEvent `json:"events"`
// }

// func CreatEventsPost(ctx *gin.Context) {
// 	type req CreatEventsPostRequest
// 	var reqData CreatEventsPostRequest
// 	var uploadedInfo filestorage.UploadInfo

// 	data := GetPayloadFromContext(ctx)
// 	file, m, err := ctx.Request.FormFile("cover_image")

// 	if err != nil {
// 		restErr := errors.NewBadRequestError("Please attach image to the request")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return
// 	}

// 	postData := req{
// 		Title:       ctx.PostForm("title"),
// 		SubTitle:    ctx.PostForm("sub_title"),
// 		Content:     ctx.PostForm("content"),
// 		ScheduledOn: ctx.PostForm("scheduled_on"),
// 	}
// 	reqData = CreatEventsPostRequest(postData)
// 	fileContentType := m.Header["Content-Type"][0]

// 	uploadFile, err := services.MinioService.UploadFile(m.Filename, file, m.Size, fileContentType)
// 	if err != nil {
// 		restErr := errors.NewBadRequestError("could not upload image to server")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return

// 	}
// 	log.Println(data)
// 	user, err := services.UsersService.GetUserByEmail(data.Username)
// 	if err != nil {

// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	uploadedInfo = uploadFile

// 	dt := date_utils.StringToDate(reqData.ScheduledOn)
// 	log.Println(time.Now().UTC())
// 	log.Println(dt)
// 	eventsData := models.ChurchEvent{
// 		OrganizerID: user.ID,
// 		CoverImage:  uploadedInfo.Key,
// 		Title:       reqData.Title,
// 		SubTitle:    reqData.SubTitle,
// 		Content:     reqData.Content,
// 		ScheduledOn: dt,
// 		ChurchJobs:  nil,
// 	}

// 	events, errr := services.EventsService.CreateEvent(eventsData)
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing create events post request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, events)

// }

// func UpdateEventsPost(ctx *gin.Context) {
// 	type req CreatEventsPostRequest
// 	var reqData CreatEventsPostRequest
// 	var uploadedInfo filestorage.UploadInfo

// 	data := GetPayloadFromContext(ctx)
// 	id := ctx.Query("id")
// 	value, _ := strconv.ParseInt(id, 10, 32)
// 	if id == "" || value == 0 {
// 		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	i, err := strconv.ParseUint(id, 10, 32)
// 	if err != nil {
// 		data := errors.NewBadRequestError("Provide an id to the request.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return

// 	}
// 	file, m, err := ctx.Request.FormFile("cover_image")

// 	if err != nil {
// 		restErr := errors.NewBadRequestError("Please attach image to the request")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return
// 	}

// 	postData := req{
// 		Title:    ctx.PostForm("title"),
// 		SubTitle: ctx.PostForm("sub_title"),
// 		Content:  ctx.PostForm("content"),
// 	}
// 	reqData = CreatEventsPostRequest(postData)
// 	fileContentType := m.Header["Content-Type"][0]

// 	uploadFile, err := services.MinioService.UploadFile(m.Filename, file, m.Size, fileContentType)
// 	if err != nil {
// 		restErr := errors.NewBadRequestError("could not upload image to server")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return

// 	}
// 	log.Println(data)
// 	uploadedInfo = uploadFile
// 	eventsData := models.ChurchEvent{
// 		CoverImage: uploadedInfo.Key,
// 		Title:      reqData.Title,
// 		SubTitle:   reqData.SubTitle,
// 		Content:    reqData.Content,
// 	}
// 	errr := services.EventsService.UpdateEventsPost(uint(i), eventsData)
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing create events post request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "events model updated",
// 	})

// }

// func GetAllEvents(ctx *gin.Context) {
// 	cacheData, errr := services.CacheService.GetEventsList(context.Background(), "events-list")

// 	if errr == nil {
// 		ctx.JSON(http.StatusOK, cacheData)
// 		return
// 	}
// 	events, count, err := services.EventsService.GetAllEvents()
// 	if err != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	// _ = services.CacheService.SetEventsList(context.Background(), events)

// 	data := GetAllEventsResponse{
// 		Total:  count,
// 		Events: events,
// 	}
// 	ctx.JSON(http.StatusOK, data)

// }

// func GetAllEventsByAuthor(ctx *gin.Context) {
// 	//id := ctx.Query("id")
// 	//value, _ := strconv.ParseInt(id, 10, 32)
// 	//if id == "" || value == 0 {
// 	//	data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
// 	//	ctx.JSON(data.Status, data)
// 	//	ctx.Abort()
// 	//	return
// 	//}
// 	//i, err := strconv.ParseUint(id, 10, 32)
// 	//if err != nil {
// 	//	data := errors.NewBadRequestError("Provide an id to the request.")
// 	//	ctx.JSON(data.Status, data)
// 	//	ctx.Abort()
// 	//	return
// 	//
// 	//}
// 	user := GetPayloadFromContext(ctx)

// 	events, count, errr := services.EventsService.GetAllEventsByAuthor(user.ID)
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	type GetAllEventsResponse2 struct {
// 		Total  int64                 `json:"total"`
// 		Events []*models.ChurchEvent `json:"events"`
// 	}
// 	data := GetAllEventsResponse2{
// 		Total:  count,
// 		Events: events,
// 	}
// 	ctx.JSON(http.StatusOK, data)

// }
// func GetAllEventsByAuthorAdmin(ctx *gin.Context) {
// 	id := ctx.Query("id")
// 	value, _ := strconv.ParseInt(id, 10, 32)
// 	if id == "" || value == 0 {
// 		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	i, err := strconv.ParseUint(id, 10, 32)
// 	if err != nil {
// 		data := errors.NewBadRequestError("Provide an id to the request.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return

// 	}
// 	events, count, errr := services.EventsService.GetAllEventsByAuthor(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	type GetAllEventsResponse2 struct {
// 		Total  int64                 `json:"total"`
// 		Events []*models.ChurchEvent `json:"events"`
// 	}
// 	data := GetAllEventsResponse2{
// 		Total:  count,
// 		Events: events,
// 	}
// 	ctx.JSON(http.StatusOK, data)

// }

// func DeleteEvent(ctx *gin.Context) {
// 	id := ctx.Query("id")
// 	value, _ := strconv.ParseInt(id, 10, 32)
// 	if id == "" || value == 0 {
// 		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	i, err := strconv.ParseUint(id, 10, 32)
// 	if err != nil {
// 		data := errors.NewBadRequestError("Provide an id to the request.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return

// 	}
// 	errr := services.EventsService.DeleteEvent(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"Message": "Successfully deleted events",
// 	})

// }

// func GetSingleEvent(ctx *gin.Context) {
// 	id := ctx.Query("id")
// 	value, _ := strconv.ParseInt(id, 10, 32)
// 	if id == "" || value == 0 {
// 		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	i, err := strconv.ParseUint(id, 10, 32)
// 	if err != nil {
// 		data := errors.NewBadRequestError("Provide an id to the request.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return

// 	}

// 	eventsData, bg := services.CacheService.GetEvents(context.Background(), "single-events")
// 	if bg == nil {
// 		log.Println(eventsData)
// 		ctx.JSON(http.StatusOK, eventsData)
// 		return
// 	}

// 	events, errr := services.EventsService.GetSingleEvent(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	_ = services.CacheService.SetEvents(context.Background(), events)
// 	ctx.JSON(http.StatusOK, events)

// }
