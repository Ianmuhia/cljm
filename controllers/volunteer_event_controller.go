package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //nolint:goimports
	"maranatha_web/services"
	"maranatha_web/utils/errors" //nolint:goimports
)

type SubScribeToEventJobRequest struct {
	Job string `json:"job"  binding:"required"`
}

//type GetAllSermonsResponse struct {
//	Total   int64           `json:"total"`
//	Sermons []models.Sermon `json:"sermons"`
//}

func SubscribeToEventJob(ctx *gin.Context) {
	var req SubScribeToEventJobRequest
	data := GetPayloadFromContext(ctx)

	if err := ctx.ShouldBindJSON(&req); err != nil {

		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}

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
	user, err := services.UsersService.GetUserByEmail(data.Username)
	if err != nil {
		data := errors.NewBadRequestError("Could not get the provided user.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	log.Println(user, i)
	//event, errr := services.EventsService.GetSingleEvent(uint(i))
	//if errr != nil {
	//	data := errors.NewBadRequestError("Could not get the provided event.")
	//	ctx.JSON(data.Status, data)
	//	ctx.Abort()
	//	return
	//}
	//job, err :=
	//services.EventsService.
	//	value := models.VolunteerChurchJob{
	//	VolunteerID: user.ID,
	//	ChurchJobID: 0,
	//}
	//partner, errr := services.VolunteerChurchJobService.CreateSubscribeToChurchJob(value)
	//if errr != nil {
	//	data := errors.NewBadRequestError("Error Processing create sermon request")
	//	ctx.JSON(data.Status, data)
	//	ctx.Abort()
	//	return
	//}

	ctx.JSON(http.StatusCreated, "ok")

}
