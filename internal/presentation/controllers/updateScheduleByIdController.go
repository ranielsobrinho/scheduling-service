package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
)

type UpdateScheduleByIdController struct {
	updateScheduleByIdUseCase usecases.UpdateSchedulingByIdUseCase
}

func NewUpdateScheduleByIdController(updateScheduleByIdUseCase usecases.UpdateSchedulingByIdUseCase) UpdateScheduleByIdController {
	return UpdateScheduleByIdController{updateScheduleByIdUseCase: updateScheduleByIdUseCase}
}

func (updateScheduleByIdController *UpdateScheduleByIdController) UpdateScheduleById(ctx *gin.Context) {
	scheduleId := ctx.Param("scheduleId")
	var scheduleModel models.SchedulingModel
	err := ctx.BindJSON(&scheduleModel)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedSchedule, err := updateScheduleByIdController.updateScheduleByIdUseCase.UpdateScheduleById(scheduleId, scheduleModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, updatedSchedule)
}
