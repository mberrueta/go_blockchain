package main

// Grpc tuto
// https://www.youtube.com/watch?v=gju-bml4kdw
import (
	"fmt"
	"log"
	"net"
	// "net/http"

	"github.com/mberrueta/test1/blockchain"
	"github.com/mberrueta/test1/proto"
	"github.com/mberrueta/test1/server"
	// "golang.org/x/net/context"
	"google.golang.org/grpc"
	// "github.com/grpc-ecosystem/go-grpc-prometheus"
	// pb "github.com/grpc-ecosystem/go-grpc-prometheus/examples/grpc-server-with-prometheus/protobuf"
	// "github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	fmt.Printf("listen port: %v", "8080")

	if err != nil {
		log.Fatalf("unable to listen port 8080: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &server.Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}
