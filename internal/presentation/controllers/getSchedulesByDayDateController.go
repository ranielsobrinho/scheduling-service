package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
)

type GetSchedulesByDayDateController struct {
	getSchedulesByDayDateUseCase usecases.GetSchedulesByDayDateUseCase
}

func NewGetSchedulesByDayDateController(getSchedulesByDayDateUseCase usecases.GetSchedulesByDayDateUseCase) GetSchedulesByDayDateController {
	return GetSchedulesByDayDateController{getSchedulesByDayDateUseCase: getSchedulesByDayDateUseCase}
}

func (getSchedulesByDayDateController *GetSchedulesByDayDateController) GetSchedulesByDayDate(ctx *gin.Context) {
	dayDate := ctx.Param("dayDate")

	if dayDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Day cannot be null",
		})
		return
	}

	schedules, err := getSchedulesByDayDateController.getSchedulesByDayDateUseCase.GetSchedulesByDayDate(dayDate)

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
