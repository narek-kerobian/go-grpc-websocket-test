package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/hashicorp/go-memdb"
	"github.com/narek-kerobian/go-grpc-websocket-test/entity"
	pb "github.com/narek-kerobian/go-grpc-websocket-test/pb/balance"
	"google.golang.org/grpc"
)

var wallet entity.Wallet

// Define new server struct interfacing the generated one
type server struct {
    pb.BalanceServiceServer
    db *memdb.MemDB
}

// Balance retrieval method
func (s *server) Balance(ctx context.Context, req *pb.BalanceServiceRequest) (*pb.BalanceServiceReply, error) {
    log.Println("[RPC Log] Received Balance request")

    userId, err := strconv.ParseUint(req.UserId, 10, 64)
    if err != nil {
        log.Fatal("[RPC Log] Invalid input")
    }

    wallet := wallet.FindOneByUserId(s.db, uint(userId))

    log.Println("[RPC Log] Responding to Balance request")

    fmt.Printf("%#v", &pb.BalanceServiceReply{
            Balance: wallet.Balance,
            })

	return &pb.BalanceServiceReply{
		Balance: wallet.Balance,
	}, nil
}

// Initialize RPC server
func InitRpc(db *memdb.MemDB) {
    // Hardcoded port is required for reverse proxy
    listener, err := net.Listen("tcp", ":10081")
    if err != nil {
        panic(err)
    }
    log.Println("[RPC Log] Listening and serving on port 10081")

    s := grpc.NewServer()
    pb.RegisterBalanceServiceServer(s, &server{db: db})
    if err := s.Serve(listener); err != nil {
        log.Fatalf("[RPC Log] Failed to serve: %v", err)
    }
}
