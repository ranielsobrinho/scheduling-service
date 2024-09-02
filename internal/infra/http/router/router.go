package router

import (
	"github.com/gin-gonic/gin"
	db "github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/helpers"
)

func Initialize() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	initializeRoutes(server, dbConnection)

	server.Run(":5050")
}
