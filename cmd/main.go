package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranielsobrinho/scheduling-service-api/internal/data/usecases"
	db "github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/helpers"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
	"github.com/ranielsobrinho/scheduling-service-api/internal/presentation/controllers"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	// Repository
	SchedulingRepository := repositories.NewSchedulingRepository(dbConnection)

	// UseCase
	SchedulingUseCase := usecases.NewGetSchedulesUseCase(SchedulingRepository)

	// Controller
	SchedulingController := controllers.NewGetSchedulesController(SchedulingUseCase)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.GET("/api/schedules", SchedulingController.GetSchedules)

	server.Run(":5050")
}
