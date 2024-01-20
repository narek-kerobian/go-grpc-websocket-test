package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narek-kerobian/go-grpc-websocket-test/config"
	"github.com/narek-kerobian/go-grpc-websocket-test/service"
)

const appPort = "10080"

func main() {
    // Initialize database
    db := service.InitCache()

    // Load mock data
    service.LoadMockData(db)

    // Initialize routes
    r := gin.Default()
    config.InitRoutes(r, db)

    r.Run(":" + appPort)
}
