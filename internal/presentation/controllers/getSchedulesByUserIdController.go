package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
)

type GetSchedulesByUserIdController struct {
	getSchedulesByUserIdUseCase usecases.GetSchedulesByUserIdUseCase
}

func NewGetSchedulesByUserIdController(getSchedulesByUserIdUseCase usecases.GetSchedulesByUserIdUseCase) GetSchedulesByUserIdController {
	return GetSchedulesByUserIdController{getSchedulesByUserIdUseCase: getSchedulesByUserIdUseCase}
}

func (getSchedulesByUserIdController *GetSchedulesByUserIdController) GetSchedulesByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	userIdNumber, err := strconv.Atoi(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	schedules, err := getSchedulesByUserIdController.getSchedulesByUserIdUseCase.GetSchedulesByUserId(userIdNumber)

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
