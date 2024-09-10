package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
)

type CreateSchedulingController struct {
	createSchedulingUseCase usecases.CreateScheduleUseCase
}

func NewCreateSchedulingController(createSchedulingUseCase usecases.CreateScheduleUseCase) CreateSchedulingController {
	return CreateSchedulingController{createSchedulingUseCase: createSchedulingUseCase}
}

func (createSchedulingController *CreateSchedulingController) CreateSchedule(ctx *gin.Context) {
	var schedulingModel models.SchedulingModel

	err := ctx.BindJSON(&schedulingModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedSchedulingId, err := createSchedulingController.createSchedulingUseCase.CreateScheduleUseCase(schedulingModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedSchedulingId)
}
