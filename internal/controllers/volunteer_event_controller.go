package controllers

// import (
// 	"log"
// 	"net/http"
// 	"strconv"
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

// 	"github.com/gin-gonic/gin" //nolint:goimports

// 	"maranatha_web/models"
// 	"maranatha_web/services"
// 	"maranatha_web/utils/errors" //nolint:goimports
// )

// type SubScribeToEventJobRequest struct {
// 	Job int `json:"job"  binding:"required"`
// }

// func SubscribeToEventJob(ctx *gin.Context) {
// 	//var req SubScribeToEventJobRequest
// 	data := GetPayloadFromContext(ctx)
// 	//if err := ctx.ShouldBindJSON(&req); err != nil {
// 	//
// 	//	restErr := errors.NewBadRequestError("invalid json body")
// 	//	ctx.JSON(restErr.Status, restErr)
// 	//	ctx.Abort()
// 	//	return
// 	//}

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
// 	user, err := services.UsersService.GetUserByEmail(data.Username)
// 	if err != nil {
// 		data := errors.NewBadRequestError("Could not get the provided user.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	job, errr := services.JobsService.GetSingleJob(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Could not get job matching that id.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	log.Println(user, i)
// 	subJob := models.VolunteerChurchJob{
// 		VolunteerID: user.ID,
// 		ChurchJobID: job.ID,
// 	}
// 	jobSubscribe, errr := services.VolunteerChurchJobService.CreateSubscribeToChurchJob(subJob)
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing create sermon request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, jobSubscribe)

// }
