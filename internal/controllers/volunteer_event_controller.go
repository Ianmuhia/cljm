package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"maranatha_web/internal/models"
	"maranatha_web/internal/utils/errors"
	"net/http"
	"strconv"
	"time"
)

type SubScribeToEventJobRequest struct {
	Job int `json:"job"  binding:"required"`
}

func (r *Repository) SubscribeToEventJob(ctx *gin.Context) {
	//var req SubScribeToEventJobRequest
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
	user, err := r.userServices.GetUserByEmail(data.Username)
	if err != nil {
		data := errors.NewBadRequestError("Could not get the provided user.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	job, errr := r.jobService.GetSingleJob(uint(i))
	if errr != nil {
		data := errors.NewBadRequestError("Could not get job matching that id.")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}
	log.Println(user, i)
	subJob := models.VolunteerChurchJob{
		VolunteerID: user.ID,
		ChurchJobID: job.ID,
	}
	jobSubscribe, err := r.volunteerService.CreateSubscribeToChurchJob(subJob)
	if errr != nil {
		data := errors.NewBadRequestError("Error Processing create sermon request")
		ctx.JSON(data.Status, data)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, jobSubscribe)

}

type UserVolunteeredJobsResponse struct {
	Total int                          `json:"total"`
	Jobs  []*models.VolunteerChurchJob `json:"jobs"`
}

func (r *Repository) GetUserVolunteeredJobs(ctx *gin.Context) {
	payload := r.GetPayloadFromContext(ctx)
	total, jobs, err := r.volunteerService.GetUserVolunteeredJobs(int(payload.ID))
	if err != nil {
		fmt.Println(err)
		resp := errors.NewBadRequestError("Could not get user volunteered jobs.")
		ctx.JSON(resp.Status, resp)
		return
	}
	response := UserVolunteeredJobsResponse{
		Total: total,
		Jobs:  jobs,
	}
	resp := SuccessResponse{
		TimeStamp: time.Now(),
		Message:   "Successfully got  user volunteered jobs",
		Status:    http.StatusOK,
		Data:      response,
	}

	ctx.JSON(resp.Status, resp)

}
