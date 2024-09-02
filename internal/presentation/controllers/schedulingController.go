package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
)

type GetSchedulesController struct {
	schedulingUseCase usecases.GetSchedulesUseCase
}

func NewGetSchedulesController(schedulingUseCase usecases.GetSchedulesUseCase) GetSchedulesController {
	return GetSchedulesController{schedulingUseCase: schedulingUseCase}
}

func (getschedulesController *GetSchedulesController) GetSchedules(ctx *gin.Context) {
	schedules, err := getschedulesController.schedulingUseCase.GetSchedules()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(schedules) == 0 {
		ctx.JSON(http.StatusOK, []models.SchedulingModel{})
		return
	}

	ctx.JSON(http.StatusOK, schedules)
}
