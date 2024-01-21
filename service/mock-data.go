package service

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/hashicorp/go-memdb"
	"github.com/narek-kerobian/go-grpc-websocket-test/entity"
)

func LoadMockData(db *memdb.MemDB) {
    log.Println("[Application Log] Loading mock data")

    for id := 1; id <= 5; id++ {
        wallet := entity.Wallet{
            Id: uint(id),
            Balance: rand.Float64() * 1000,
        }
        wallet.Insert(db);

        user := entity.User{
            Id: uint(id),
            Name: fmt.Sprintf("%s: %d", "User", rand.Int()),
            WalletId: wallet.Id,
        }
        user.Insert(db);
    }
}
