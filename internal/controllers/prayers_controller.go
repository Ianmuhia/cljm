package controllers

import (
	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePrayerRequest struct {
	Content string `json:"content"`
}

type GetAllPrayerRequestResponse struct {
	Total  int64            `json:"total"`
	Prayer []*models.Prayer `json:"prayer"`
}

func (r *Repository) CreatPrayerRequest(ctx *gin.Context) {
	var req CreatePrayerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	data := r.GetPayloadFromContext(ctx)
	user, err := r.userServices.GetUserByEmail(data.Username)

	if err != nil {
		data := errors.NewBadRequestError("Error Processing create prayer request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	postData := models.Prayer{

		AuthorID: user.ID,
		Content:  req.Content,
	}

	prayer, errr := r.prayerRequestService.CreatePrayerRequest(postData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create prayer request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, prayer)

}

func (r *Repository) UpdatePrayerRequest(ctx *gin.Context) {

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

	var req CreatePrayerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	prayerData := models.Prayer{
		Content: req.Content,
	}
	errr := r.prayerRequestService.UpdatePrayerRequest(uint(i), prayerData)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create prayer post request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "prayer model updated",
	})

}

func (r *Repository) GetAllPrayerRequests(ctx *gin.Context) {
	//cacheData, errr := r..GetPrayerList(context.Background(), "prayers-list")
	//
	//if errr == nil {
	//	ctx.JSON(http.StatusOK, cacheData)
	//	return
	//}
	prayer, count, err := r.prayerRequestService.GetAllPrayerRequests()
	if err != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	data := GetAllPrayerRequestResponse{
		Total:  count,
		Prayer: prayer,
	}
	ctx.JSON(http.StatusOK, data)

}

func (r *Repository) GetAllPrayerRequestsByAuthor(ctx *gin.Context) {
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
	prayer, count, errr := r.prayerRequestService.GetAllPrayerRequestsByAuthor(user.ID)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	type GetAllPrayerResponse2 struct {
		Total  int64            `json:"total"`
		Prayer []*models.Prayer `json:"prayer"`
	}
	data := GetAllPrayerResponse2{
		Total:  count,
		Prayer: prayer,
	}
	ctx.JSON(http.StatusOK, data)

}
func (r *Repository) GetAllPrayerRequestsByAuthorAdmin(ctx *gin.Context) {
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
	prayer, count, errr := r.prayerRequestService.GetAllPrayerRequestsByAuthor(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	type GetAllPrayerResponse2 struct {
		Total  int64            `json:"total"`
		Prayer []*models.Prayer `json:"prayer"`
	}
	data := GetAllPrayerResponse2{
		Total:  count,
		Prayer: prayer,
	}
	ctx.JSON(http.StatusOK, data)

}

func (r *Repository) DeletePrayerRequest(ctx *gin.Context) {
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
	errr := r.prayerRequestService.DeletePrayerRequest(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted prayer",
	})

}

func (r *Repository) GetSinglePrayerRequest(ctx *gin.Context) {

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

	//prayerData, bg := services.CacheService.GetPrayer(context.Background(), "single-prayer")
	//if bg == nil {
	//	log.Println(prayerData)
	//	ctx.JSON(http.StatusOK, prayerData)
	//	return
	//}

	prayer, errr := r.prayerRequestService.GetSinglePrayerRequest(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	//_ = services.CacheService.SetPrayer(context.Background(), prayer)
	ctx.JSON(http.StatusOK, prayer)

}
