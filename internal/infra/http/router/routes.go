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
	DeleteScheduleByIdUseCase := usecases.NewDeleteScheduleByIdUseCase(SchedulingRepository)
	GetScheduleByIdUseCase := usecases.NewGetScheduleByIdUseCase(SchedulingRepository)
	UpdateScheduleByIdUseCase := usecases.NewUpdateSchedulingByIdUseCase(SchedulingRepository)
	GetSchedulesByDayDateUseCase := usecases.NewGetSchedulesByDayDateUseCase(SchedulingRepository)
	GetSchedulesByMonthUseCase := usecases.NewGetSchedulesByMonthUseCase(SchedulingRepository)

	// Controllers
	GetSchedulingController := controllers.NewGetSchedulesController(GetSchedulingUseCase)
	CreateSchedulingController := controllers.NewCreateSchedulingController(CreateSchedulingUseCase)
	GetSchedulesByUserIdController := controllers.NewGetSchedulesByUserIdController(GetSchedulesByUserIdUseCase)
	DeleteScheduleByIdController := controllers.NewDeleteScheduleByIdController(DeleteScheduleByIdUseCase)
	GetScheduleByIdController := controllers.NewGetScheduleByIdController(GetScheduleByIdUseCase)
	UpdateScheduleByIdController := controllers.NewUpdateScheduleByIdController(UpdateScheduleByIdUseCase)
	GetSchedulesByDayDateController := controllers.NewGetSchedulesByDayDateController(GetSchedulesByDayDateUseCase)
	GetSchedulesByMonthController := controllers.NewGetSchedulesByMonthController(GetSchedulesByMonthUseCase)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/schedules", AuthMiddleware(), GetSchedulingController.GetSchedules)
		v1.POST("/schedules", AuthMiddleware(), CreateSchedulingController.CreateSchedule)
		v1.GET("/schedules/:userId", AuthMiddleware(), GetSchedulesByUserIdController.GetSchedulesByUserId)
		v1.DELETE("/schedules/:scheduleId", AuthMiddleware(), DeleteScheduleByIdController.DeleteScheduleById)
		v1.GET("/schedule/:scheduleId", AuthMiddleware(), GetScheduleByIdController.GetScheduleById)
		v1.PUT("/schedules/:scheduleId", AuthMiddleware(), UpdateScheduleByIdController.UpdateScheduleById)
		v1.GET("/schedules/get-by-date/:dayDate", AuthMiddleware(), GetSchedulesByDayDateController.GetSchedulesByDayDate)
		v1.GET("/schedules/get-by-month/:dayMonth", AuthMiddleware(), GetSchedulesByMonthController.GetSchedulesByMonth)
	}
}
