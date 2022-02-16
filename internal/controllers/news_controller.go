package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"

	"maranatha_web/internal/controllers/token"
	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
)

type CreatNewsPostRequest struct {
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"sub_title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type GetAllNewsResponse struct {
	Total int            `json:"total"`
	News  []*models.News `json:"news"`
}

//TODO:Update user details
func (r *Repository) CreatNewsPost(ctx *gin.Context) {
	type req CreatNewsPostRequest
	var reqData CreatNewsPostRequest
	var uploadedInfo minio.UploadInfo

	data := r.GetPayloadFromContext(ctx)
	file, m, err := ctx.Request.FormFile("cover_image")
	// form, _ := ctx.MultipartForm()
	// files := form.File["other_images"]

	// for _, file := range files {
	// 	log.Println(file.Filename)

	// 	// Upload the file to specific dst.
	// 	// c.SaveUploadedFile(file, dst)
	// }
	log.Println(file)
	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	postData := req{
		Title:    ctx.PostForm("title"),
		SubTitle: ctx.PostForm("sub_title"),
		Content:  ctx.PostForm("content"),
	}
	reqData = CreatNewsPostRequest(postData)

	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	uploadedInfo = uploadFile
	value := models.News{
		AuthorID:   data.ID,
		CoverImage: uploadedInfo.Key,
		Title:      reqData.Title,
		SubTitle:   reqData.SubTitle,
		Content:    reqData.Content,
	}
	news, err := r.newsService.CreateNewsPost(value)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing create news post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	resp := NewStatusCreatedResponse("News created successfully", news)

	ctx.JSON(resp.Status, resp)

}

func (r *Repository) UpdateNewsPost(ctx *gin.Context) {
	type req CreatNewsPostRequest
	var reqData CreatNewsPostRequest
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
		Title:    ctx.PostForm("title"),
		SubTitle: ctx.PostForm("sub_title"),
		Content:  ctx.PostForm("content"),
	}
	reqData = CreatNewsPostRequest(postData)
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
	newsData := models.News{
		CoverImage: uploadedInfo.Key,
		Title:      reqData.Title,
		SubTitle:   reqData.SubTitle,
		Content:    reqData.Content,
	}
	errr := r.newsService.UpdateNewsPost(uint(i), newsData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create news post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "news model updated",
	})

}

func (r *Repository) GetAllNewsPost(ctx *gin.Context) {
	//cacheData, err := services.CacheService.GetNewsList(context.Background(), "news-list")
	//
	//if err == nil {
	//	ctx.JSON(http.StatusOK, cacheData)
	//	return
	//}
	data, count, err := r.newsService.GetAllNewsPost()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	// _ = services.CacheService.SetNewsList(context.Background(), news)

	news := GetAllNewsResponse{
		Total: count,
		News:  data,
	}

	resp := NewStatusOkResponse("Got news successfully", news)

	ctx.JSON(resp.Status, resp)

}

func (r *Repository) GetAllNewsPostByAuthor(ctx *gin.Context) {
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
	news, count, err := r.newsService.GetAllNewsPostByAuthor(uint(i))
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
	resp := NewStatusOkResponse("All news by this author.", data)
	ctx.JSON(resp.Status, resp)
}

func (r *Repository) DeleteNewsPost(ctx *gin.Context) {
	id := ctx.Query("id")
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		data := errors.NewBadRequestError("Provide an id to the request.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return

	}
	err = r.newsService.DeleteNewsPost(uint(i))
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	resp := NewDeleteResponse("Successfully deleted news", nil)
	ctx.JSON(resp.Status, resp)

}

func (r *Repository) GeSingleNewsPost(ctx *gin.Context) {
	id := ctx.Query("id")

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		data := errors.NewBadRequestError("Provide an id to the request.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	//newsData, bg := services.CacheService.GetNews(context.Background(), "single-news")
	//if bg == nil {
	//	log.Println(newsData)
	//	ctx.JSON(http.StatusOK, newsData)
	//	return
	//}

	news, err := r.newsService.GetSingleNewsPost(uint(i))
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
	resp := NewStatusOkResponse("Get single new success.", news)
	//_ = services.CacheService.SetNews(context.Background(), news)
	ctx.JSON(resp.Status, resp)

}

func (r *Repository) BatchUpload(ctx *gin.Context) {
	var _ minio.PutObjectOptions

	form, _ := ctx.MultipartForm()
	files := form.File["other_images"]
	log.Println(len(files))
	for _, file := range files {
		log.Println(file.Filename)
		open, err := file.Open()
		if err != nil {
			return
		}
		err = open.Close()
		if err != nil {
			log.Println(err)
			return
		}

		go func() {
			uploadFile, err := r.MinoStorage.UploadFile(file.Filename, open, file.Size, "")
			if err != nil {
				log.Println(err)
			}
			log.Println(uploadFile)
		}()
		if err != nil {
			restErr := errors.NewBadRequestError("could not upload image to server")
			ctx.JSON(restErr.Status, restErr)
			ctx.Abort()
			return

		}
		// Upload the file to specific dst.
		//err := ctx.SaveUploadedFile(file, "logs")
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
	}
}

func (r *Repository) GetPayloadFromContext(ctx *gin.Context) *token.Payload {
	payload, exists := ctx.Get("authorization_payload")
	log.Println(payload)
	if !exists {
		restErr := errors.NewBadRequestError("could not get auth_payload from context")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()

	}
	data := payload.(*token.Payload)
	_, err := r.userServices.GetUserByEmail(data.Username)
	if err != nil {
		//log.Println(user)
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()

	}
	//TODO:upload news images
	return data

}
