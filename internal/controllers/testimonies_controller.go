package controllers

import (
	"github.com/gin-gonic/gin"
	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
	"net/http"
	"strconv"
)

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"strconv"
//

// 	"github.com/gin-gonic/gin"

// 	"maranatha_web/models"
// 	"maranatha_web/services"
// 	"maranatha_web/utils/errors"
// )

type CreateTestimonyRequest struct {
	Content string `json:"content" binding:"required"`
}

type GetAllTestimoniesResponse struct {
	Total       int64                 `json:"total"`
	Testimonies []*models.Testimonies `json:"testimonies"`
}

func (r *Repository) CreateTestimony(ctx *gin.Context) {

	var req CreateTestimonyRequest
	data := r.GetPayloadFromContext(ctx)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

	user, err := r.userServices.GetUserByEmail(data.Username)
	if err != nil {
		return
	}
	postData := models.Testimonies{
		Content:  req.Content,
		AuthorID: user.ID,
	}
	testimony, errr := r.testimonyService.CreateTestimony(postData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create testimonies post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, testimony)

}

func (r *Repository) UpdateTestimony(ctx *gin.Context) {
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
	var req CreateTestimonyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	testimoniesData := models.Testimonies{
		Content: req.Content,
	}
	errr := r.testimonyService.UpdateTestimony(uint(i), testimoniesData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create genre post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "testimonies model updated",
	})

}

func (r *Repository) GetAllTestimonies(ctx *gin.Context) {
	//cacheData, errr := services.CacheService.GetTestimoniesList(context.Background(), "testimonies-list")
	//
	//if errr == nil {
	//	ctx.JSON(http.StatusOK, cacheData)
	//	return
	//}
	testimonies, count, err := r.testimonyService.GetAllTestimonies()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	data := GetAllTestimoniesResponse{
		Total:       count,
		Testimonies: testimonies,
	}
	ctx.JSON(http.StatusOK, data)

}

func (r *Repository) GetAllTestimoniesByAuthor(ctx *gin.Context) {
	//id := ctx.Query("id")
	//value, _ := strconv.ParseInt(id, 10, 32)
	//if id == "" || value == 0 {
	//	data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
	//	ctx.JSON(data.Status, data)
	//	ctx.Abort()
	//	return
	//}
	//i, err := strconv.ParseUint(id, 10, 32)
	//if err != nil {
	//	data := errors.NewBadRequestError("Provide an id to the request.")
	//	ctx.JSON(data.Status, data)
	//	ctx.Abort()
	//	return
	//
	//}
	user := r.GetPayloadFromContext(ctx)

	testimonies, count, errr := r.testimonyService.GetAllTestimoniesByAuthor(user.ID)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	type GetAllTestimoniesResponse2 struct {
		Total       int64                 `json:"total"`
		Testimonies []*models.Testimonies `json:"testimonies"`
	}
	data := GetAllTestimoniesResponse2{
		Total:       count,
		Testimonies: testimonies,
	}
	ctx.JSON(http.StatusOK, data)

}
func (r *Repository) GetAllTestimoniesByAuthorAdmin(ctx *gin.Context) {
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
	testimonies, count, errr := r.testimonyService.GetAllTestimoniesByAuthor(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	type GetAllTestimoniesResponse2 struct {
		Total       int64                 `json:"total"`
		Testimonies []*models.Testimonies `json:"testimonies"`
	}
	data := GetAllTestimoniesResponse2{
		Total:       count,
		Testimonies: testimonies,
	}
	ctx.JSON(http.StatusOK, data)

}

func (r *Repository) DeleteTestimony(ctx *gin.Context) {
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
	errr := r.testimonyService.DeleteTestimony(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted testimonies",
	})

}

func (r *Repository) GetSingleTestimony(ctx *gin.Context) {

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

	//testimoniesData, bg := services.CacheService.GetTestimonies(context.Background(), "single-testimonies")
	//if bg == nil {
	//	log.Println(testimoniesData)
	//	ctx.JSON(http.StatusOK, testimoniesData)
	//	return
	//}

	testimonies, errr := r.testimonyService.GetSingleTestimony(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	//_ = services.CacheService.SetTestimonies(context.Background(), testimonies)
	ctx.JSON(http.StatusOK, testimonies)

}
