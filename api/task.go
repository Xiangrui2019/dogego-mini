package api

import (
	"dogego-mini/services"

	"github.com/gin-gonic/gin"
)

func TestAsyncTask(context *gin.Context) {
	service := services.TestAsyncTaskService{}

	resp := service.Test()

	context.JSON(resp.Code, resp)
}
