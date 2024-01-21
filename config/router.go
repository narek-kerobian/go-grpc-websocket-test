package config

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
	"github.com/narek-kerobian/go-grpc-websocket-test/controller/api"
	"github.com/narek-kerobian/go-grpc-websocket-test/controller/ws"
	"github.com/narek-kerobian/go-grpc-websocket-test/middleware"
)

func InitRoutes(r *gin.Engine, db *memdb.MemDB) {

    // Load middlewares
    r.Use(middleware.ExposeGinEngine(r))

    // API routes
    groupApi := r.Group("/api")
    {
        groupWallet := groupApi.Group("/wallet")
        {
            groupWallet.POST("/deposit", api.WalletDeposit(db))
            groupWallet.POST("/withdraw", api.WalletWithdraw(db))
            groupWallet.GET("/balance/:UserId", api.WalletBalance(db))
        }
    }

    // Websocket routes
    groupWs := r.Group("/ws")
    {
        groupWs.GET("/update", ws.HandleSocketConnections)
    }

}

