package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
	"github.com/ranielsobrinho/scheduling-service-api/internal/presentation/controllers"
)

func initializeRoutes(router *gin.Engine, dbConnection *sql.DB) {

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Repository
	SchedulingRepository := repositories.NewSchedulingRepository(dbConnection)

	// UseCases
	GetSchedulingUseCase := usecases.NewGetSchedulesUseCase(SchedulingRepository)
	CreateSchedulingUseCase := usecases.NewCreateScheduleUseCase(SchedulingRepository)
	GetSchedulesByUserIdUseCase := usecases.NewGetSchedulesByUserIdUseCase(SchedulingRepository)

	// Controllers
	GetSchedulingController := controllers.NewGetSchedulesController(GetSchedulingUseCase)
	CreateSchedulingController := controllers.NewCreateSchedulingController(CreateSchedulingUseCase)
	GetSchedulesByUserIdController := controllers.NewGetSchedulesByUserIdController(GetSchedulesByUserIdUseCase)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/schedules", GetSchedulingController.GetSchedules)
		v1.POST("/schedules", CreateSchedulingController.CreateSchedule)
		v1.GET("/schedules/:userId", GetSchedulesByUserIdController.GetSchedulesByUserId)
	}
}
