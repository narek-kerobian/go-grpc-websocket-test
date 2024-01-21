package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
	"github.com/narek-kerobian/go-grpc-websocket-test/entity"
)

var (
    errors []string
    wallet entity.Wallet
    balanceMovementRequestObject entity.BalanceMovementRequestObject
    balanceRouteParamObject entity.BalanceParamParamObject
)

func WalletDeposit(db *memdb.MemDB) gin.HandlerFunc {
    return func (c *gin.Context) {
        if err := c.Bind(&balanceMovementRequestObject); err != nil {
            errors = append(errors, "Invalid input: " + err.Error())
            c.JSON(400, wallet.BuildResponse(errors))
            return
        }

        // Update balance
        if _, err := wallet.AugmentBalance(db, balanceMovementRequestObject); err != nil {
            errors = append(errors, "Invalid input: " + err.Error())
            c.JSON(400, wallet.BuildResponse(errors))
            errors = []string{}
            return
        }

        // Build and return the response
        c.JSON(200, wallet.BuildResponse(errors))
    }
}

func WalletWithdraw(db *memdb.MemDB) gin.HandlerFunc {
    return func (c *gin.Context) {
        // Bind route param to balanceRouteParamObject
        if err := c.Bind(&balanceMovementRequestObject); err != nil {
            errors = append(errors, "Invalid input: " + err.Error())
            c.JSON(400, wallet.BuildResponse(errors))
            errors = []string{}
            return
        }

        // Fake balance
        if _, err := wallet.DeductBalance(db, balanceMovementRequestObject); err != nil {
            errors = append(errors, "Invalid input: " + err.Error())
            c.JSON(400, wallet.BuildResponse(errors))
            errors = []string{}
            return
        }

        // Build and return the response
        c.JSON(200, wallet.BuildResponse(errors))
    }
}

func WalletBalance(db *memdb.MemDB) gin.HandlerFunc {
    return func (c *gin.Context) {
        // Bind route param to balanceRouteParamObject
        if err := c.ShouldBindUri(&balanceRouteParamObject); err != nil {
            errors = append(errors, "Invalid input: " + err.Error())
            c.JSON(400, wallet.BuildResponse(errors))
            errors = []string{}
            return
        }

        // Get wallet balance by userId
        wallet = *wallet.FindOneByUserId(db, balanceRouteParamObject.UserId)

        // Build and return the response
        c.JSON(200, wallet.BuildResponse(errors))
    }
}

