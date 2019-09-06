package services

import (
	"dogego-mini/serializer"
	"dogego-mini/tasks"
	"dogego-mini/utils"
	"net/http"
)

type TestAsyncTaskService struct {
}

func (service *TestAsyncTaskService) Test() *serializer.Response {
	err := utils.RunAsyncTask(tasks.TimeTask1, "TH")

	if err != nil {
		return utils.BuildErrorResponse(err)
	}

	return &serializer.Response{
		Code:    http.StatusOK,
		Message: "Successfly send task async task.",
	}
}
