package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narek-kerobian/go-grpc-websocket-test/config"
	"github.com/narek-kerobian/go-grpc-websocket-test/controller/rpc"
	"github.com/narek-kerobian/go-grpc-websocket-test/service"
)

func main() {
    // Initialize database
    db := service.InitCache()

    // Load mock data
    service.LoadMockData(db)

    // Initialize routes
    r := gin.Default()
    config.InitRoutes(r, db)

    // Start the RPC server
    go rpc.InitRpc(db)

    // Start the http server
    // Hardcoded port is required for the reverse proxy
    r.Run(":10080")
}
