package controllers

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"strconv"
//

// 	"github.com/gin-gonic/gin" //nolint:goimports
// 	"github.com/filestorage/filestorage-go/v7"

// 	"maranatha_web/controllers/token"
// 	"maranatha_web/models"
// 	"maranatha_web/services"
// 	"maranatha_web/utils/errors" //nolint:goimports
// )

// type CreatNewsPostRequest struct {
// 	Title    string `json:"title" binding:"required"`
// 	SubTitle string `json:"sub_title" binding:"required"`
// 	Content  string `json:"content" binding:"required"`
// }

// type GetAllNewsResponse struct {
// 	Total int64          `json:"total"`
// 	News  []*models.News `json:"news"`
// }

// func CreatNewsPost(ctx *gin.Context) {
// 	type req CreatNewsPostRequest
// 	var reqData CreatNewsPostRequest
// 	var uploadedInfo filestorage.UploadInfo

// 	data := GetPayloadFromContext(ctx)
// 	file, m, err := ctx.Request.FormFile("cover_image")
// 	form, _ := ctx.MultipartForm()
// 	files := form.File["other_images"]

// 	for _, file := range files {
// 		log.Println(file.Filename)

// 		// Upload the file to specific dst.
// 		// c.SaveUploadedFile(file, dst)
// 	}

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
// 	reqData = CreatNewsPostRequest(postData)

// 	//TODO: Rework this.
// 	user, err := services.UsersService.GetUserByEmail(data.Username)
// 	if err != nil {

// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	fileContentType := m.Header["Content-Type"][0]

// 	uploadFile, err := services.MinioService.UploadFile(m.Filename, file, m.Size, fileContentType)
// 	if err != nil {
// 		restErr := errors.NewBadRequestError("could not upload image to server")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return

// 	}
// 	uploadedInfo = uploadFile
// 	value := models.News{
// 		AuthorID:   user.ID,
// 		CoverImage: uploadedInfo.Key,
// 		Title:      reqData.Title,
// 		SubTitle:   reqData.SubTitle,
// 		Content:    reqData.Content,
// 	}
// 	news, errr := services.NewsService.CreateNewsPost(value)
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing create news post request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, news)

// }

// func UpdateNewsPost(ctx *gin.Context) {
// 	type req CreatNewsPostRequest
// 	var reqData CreatNewsPostRequest
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
// 	reqData = CreatNewsPostRequest(postData)
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
// 	newsData := models.News{
// 		CoverImage: uploadedInfo.Key,
// 		Title:      reqData.Title,
// 		SubTitle:   reqData.SubTitle,
// 		Content:    reqData.Content,
// 	}
// 	errr := services.NewsService.UpdateNewsPost(uint(i), newsData)
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing create news post request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "news model updated",
// 	})

// }

// func GetAllNewsPost(ctx *gin.Context) {
// 	cacheData, errr := services.CacheService.GetNewsList(context.Background(), "news-list")

// 	if errr == nil {
// 		ctx.JSON(http.StatusOK, cacheData)
// 		return
// 	}
// 	news, count, err := services.NewsService.GetAllNewsPost()
// 	if err != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	// _ = services.CacheService.SetNewsList(context.Background(), news)

// 	data := GetAllNewsResponse{
// 		Total: count,
// 		News:  news,
// 	}
// 	ctx.JSON(http.StatusOK, data)

// }

// func GetAllNewsPostByAuthor(ctx *gin.Context) {
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
// 	news, count, errr := services.NewsService.GetAllNewsPostByAuthor(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	type GetAllNewsResponse2 struct {
// 		Total int64          `json:"total"`
// 		News  []*models.News `json:"news"`
// 	}
// 	data := GetAllNewsResponse2{
// 		Total: count,
// 		News:  news,
// 	}
// 	ctx.JSON(http.StatusOK, data)

// }

// func DeleteNewsPost(ctx *gin.Context) {
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
// 	errr := services.NewsService.DeleteNewsPost(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"Message": "Successfully deleted news",
// 	})

// }

// func GeSingleNewsPost(ctx *gin.Context) {
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

// 	newsData, bg := services.CacheService.GetNews(context.Background(), "single-news")
// 	if bg == nil {
// 		log.Println(newsData)
// 		ctx.JSON(http.StatusOK, newsData)
// 		return
// 	}

// 	news, errr := services.NewsService.GetSingleNewsPost(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	_ = services.CacheService.SetNews(context.Background(), news)
// 	ctx.JSON(http.StatusOK, news)

// }

// func GetPayloadFromContext(ctx *gin.Context) *token.Payload {
// 	payload, exists := ctx.Get("authorization_payload")
// 	if !exists {
// 		restErr := errors.NewBadRequestError("could not get auth_payload from context")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()

// 	}
// 	data := payload.(*token.Payload)
// 	user, err := services.UsersService.GetUserByEmail(data.Username)
// 	if err != nil {
// 		//log.Println(user)
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()

// 	}
// 	log.Println(user)
// 	//TODO:upload news images
// 	//TODO:Work on user profile, reset password , forgot password

// 	return data

// }
