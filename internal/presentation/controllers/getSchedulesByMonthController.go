package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
)

type GetSchedulesByMonthController struct {
	getScheduleByMonthUseCase usecases.GetSchedulesByMonthUseCase
}

func NewGetSchedulesByMonthController(getSchedulesByMonthUseCase usecases.GetSchedulesByMonthUseCase) GetSchedulesByMonthController {
	return GetSchedulesByMonthController{getScheduleByMonthUseCase: getSchedulesByMonthUseCase}
}

func (getSchedulesByMonthController *GetSchedulesByMonthController) GetSchedulesByMonth(ctx *gin.Context) {
	dayMonth := ctx.Param("dayMonth")

	if dayMonth == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Month cannot be null",
		})
		return
	}

	schedules, err := getSchedulesByMonthController.getScheduleByMonthUseCase.GetSchedulesByMonth(dayMonth)

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
