package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
)

type GetScheduleByIdController struct {
	getScheduleByIdUseCase usecases.GetScheduleByIdUseCase
}

func NewGetScheduleByIdController(getScheduleByIdUseCase usecases.GetScheduleByIdUseCase) GetScheduleByIdController {
	return GetScheduleByIdController{getScheduleByIdUseCase: getScheduleByIdUseCase}
}

func (getScheduleByIdController *GetScheduleByIdController) GetScheduleById(ctx *gin.Context) {
	scheduleId := ctx.Param("scheduleId")

	if scheduleId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Schedule id cannot be null",
		})
	}

	schedule, err := getScheduleByIdController.getScheduleByIdUseCase.GetSchedulesById(scheduleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, schedule)
}
