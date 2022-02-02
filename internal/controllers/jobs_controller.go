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

// 	"github.com/gin-gonic/gin" //nolint:goimports

// 	"maranatha_web/models"
// 	"maranatha_web/services"
// 	"maranatha_web/utils/errors" //nolint:goimports
// )

// type GetAllJobsResponse struct {
// 	Total int64               `json:"total"`
// 	Jobs  []*models.ChurchJob `json:"jobs"`
// }

// type CreateEventJobRequest struct {
// 	Job     string `json:"job"`
// 	EventId int    `json:"event_id"`
// }

// type UpdateEventJobRequest struct {
// 	Job string `json:"job"`
// }

// func GetAllEventJobs(ctx *gin.Context) {
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

// 	jobs, count, errr := services.JobsService.GetAllEventJobs(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	data := GetAllJobsResponse{
// 		Total: count,
// 		Jobs:  jobs,
// 	}
// 	ctx.JSON(http.StatusOK, data)

// }

// func GetJobByEvent(ctx *gin.Context) {

// 	eventId := ctx.Query("event_id")
// 	jobId := ctx.Query("job_id")

// 	value, _ := strconv.ParseInt(eventId, 10, 32)
// 	if eventId == "" || value == 0 {
// 		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	i, err := strconv.ParseUint(eventId, 10, 32)
// 	if err != nil {
// 		data := errors.NewBadRequestError("Provide an id to the request.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return

// 	}

// 	value2, _ := strconv.ParseInt(jobId, 10, 32)
// 	if jobId == "" || value2 == 0 {
// 		data := errors.NewBadRequestError("Provide an id to the request.Id cannot be zero")
// 		log.Println("herere")

// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}
// 	i2, err := strconv.ParseUint(jobId, 10, 32)
// 	if err != nil {
// 		data := errors.NewBadRequestError("Provide an id to the request.")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return

// 	}

// 	jobs, errr := services.JobsService.GetJobByEvent(uint(i), uint(i2))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, jobs)

// }

// func CreateEventJob(ctx *gin.Context) {
// 	var req CreateEventJobRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		restErr := errors.NewBadRequestError("invalid json body")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return
// 	}

// 	event, err := services.EventsService.GetSingleEvent(uint(req.EventId))
// 	if err != nil {
// 		restErr := errors.NewBadRequestError("Error processing the request .Could not get the event by that id.")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return
// 	}

// 	data := models.ChurchJob{
// 		Duty:          req.Job,
// 		ChurchEventID: event.ID,
// 	}

// 	job, errr := services.JobsService.CreateEventJob(data)
// 	if errr != nil {
// 		restErr := errors.NewBadRequestError("Error creating the job.")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, job)

// }

// func DeleteJob(ctx *gin.Context) {
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
// 	errr := services.JobsService.DeleteJob(uint(i))
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"Message": "Successfully deleted job",
// 	})

// }
// func UpdateJob(ctx *gin.Context) {
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
// 	var req UpdateEventJobRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		restErr := errors.NewBadRequestError("invalid json body")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return
// 	}
// 	data := models.ChurchJob{
// 		Duty: req.Job,
// 	}

// 	errr := services.JobsService.UpdateJob(uint(i), data)
// 	if errr != nil {
// 		data := errors.NewBadRequestError("Error Processing request")
// 		ctx.JSON(data.Status, data)
// 		ctx.Abort()
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"Message": "Successfully updated job",
// 	})

// }
