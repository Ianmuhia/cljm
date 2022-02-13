package controllers

import (
	"net/http"
	"time"
)

type SuccessResponse struct {
	TimeStamp time.Time   `json:"time_stamp"`
	Message   string      `json:"message"`
	Status    int         `json:"status"`
	Data      interface{} `json:"data"`
}

func NewStatusOkResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusOK,
		Data:      data,
	}
}

func NewStatusCreatedResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusCreated,
		Data:      data,
	}
}

func NewDeleteResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusNoContent,
		Data:      data,
	}
}
