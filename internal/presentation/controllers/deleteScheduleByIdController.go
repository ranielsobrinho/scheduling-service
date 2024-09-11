package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
)

type DeleteScheduleByIdController struct {
	deleteScheduleByIdUseCase usecases.DeleteScheduleByIdUseCase
}

func NewDeleteScheduleByIdController(deleteScheduleByIdUseCase usecases.DeleteScheduleByIdUseCase) DeleteScheduleByIdController {
	return DeleteScheduleByIdController{deleteScheduleByIdUseCase: deleteScheduleByIdUseCase}
}

func (deleteScheduleByIdController *DeleteScheduleByIdController) DeleteScheduleById(ctx *gin.Context) {
	scheduleId := ctx.Param("scheduleId")

	if scheduleId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Schedule id cannot be null",
		})
	}

	err := deleteScheduleByIdController.deleteScheduleByIdUseCase.DeleteScheduleById(scheduleId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Schedule deleted successfully",
	})
}
