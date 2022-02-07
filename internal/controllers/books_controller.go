package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"

	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
)

type CreateBookPostRequest struct {
	Title    string `json:"title" binding:"required"`
	Synopsis string `json:"synopsis" binding:"required"`
	Genre    string `json:"genre" binding:"required"`
}

type GetAllBooksResponse struct {
	Total int64           `json:"total"`
	Books []*models.Books `json:"books"`
}

//CreateBook is incomplete, integrate with Genre.
func (r *Repository) CreateBook(ctx *gin.Context) {
	type req CreateBookPostRequest
	var reqData CreateBookPostRequest
	var uploadedInfo minio.UploadInfo

	file, m, err := ctx.Request.FormFile("file")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach file to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	postData := req{
		Title:    ctx.PostForm("title"),
		Synopsis: ctx.PostForm("synopsis"),
		Genre:    ctx.PostForm("genre"),
	}
	reqData = CreateBookPostRequest(postData)

	fileContentType := m.Header["Content-Type"][0]
	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload file to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	uploadedInfo = uploadFile
	if err != nil {

		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	creator := r.GetPayloadFromContext(ctx)
	user, err := r.userServices.GetUserByEmail(creator.Username)
	if err != nil {

		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	log.Println(user)

	genre, err := r.genreService.GetSingleGenre(postData.Genre)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data.Message)
		ctx.Abort()
		return
	}

	value := models.Books{
		Title:       reqData.Title,
		Synopsis:    reqData.Synopsis,
		CreatedBy:   user,
		CreatedByID: user.ID,
		GenreID:     genre.ID,
		File:        uploadedInfo.Key,
	}
	books, err := r.bookService.CreateBooksPost(value)
	if err != nil {
		data := errors.NewBadRequestError("Error Processing creating book post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, books)
}

func (r *Repository) UpdateBook(ctx *gin.Context) {
	type req CreateBookPostRequest
	var reqData CreateBookPostRequest
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
	file, m, err := ctx.Request.FormFile("file")

	if err != nil {
		restErr := errors.NewBadRequestError("Please attach image to the request")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	postData := req{
		Title:    ctx.PostForm("title"),
		Synopsis: ctx.PostForm("synopsis"),
		//Author:   ctx.PostForm("author"),
	}
	reqData = CreateBookPostRequest(postData)
	fileContentType := m.Header["Content-Type"][0]

	uploadFile, err := r.MinoStorage.UploadFile(m.Filename, file, m.Size, fileContentType)
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	if err != nil {
		restErr := errors.NewBadRequestError("could not upload image to server")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return

	}
	log.Println(data)
	uploadedInfo = uploadFile

	booksData := models.Books{
		Title:    reqData.Title,
		Synopsis: reqData.Synopsis,
		File:     uploadedInfo.Key,
	}
	errr := r.bookService.UpdateBooksPost(uint(i), booksData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create books post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	res := SuccessResponse{
		TimeStamp: time.Now(),
		Message:   "Book updated successfully",
		Status:    http.StatusOK,
	}
	ctx.JSON(http.StatusCreated, res)
}

func (r *Repository) GetAllBooksPost(ctx *gin.Context) {
	//cacheData, errr := services.CacheService.GetBooksList(context.Background(), "books-list")
	//
	//if errr == nil {
	//	ctx.JSON(http.StatusOK, cacheData)
	//	return
	//}
	books, count, err := r.bookService.GetAllBooks()

	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	data := GetAllBooksResponse{
		Total: count,
		Books: books,
	}

	ctx.JSON(http.StatusOK, data)

}

func (r *Repository) DeleteBook(ctx *gin.Context) {
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
	errr := r.bookService.DeleteBook(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	res := SuccessResponse{
		TimeStamp: time.Now(),
		Message:   "Book deleted successfully",
		Status:    http.StatusOK,
	}
	ctx.JSON(http.StatusCreated, res)
}

func (r *Repository) GetSingleBookPost(ctx *gin.Context) {
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

	//booksData, bg := services.CacheService.GetBooks(context.Background(), "single-books")
	//if bg == nil {
	//	log.Println(booksData)
	//
	//	ctx.JSON(http.StatusOK, booksData)
	//
	//	return
	//}

	books, errr := r.bookService.GetSingleBooksPost(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	//_ = services.CacheService.SetBooks(context.Background(), books)
	ctx.JSON(http.StatusOK, books)
}
