// Package main handle the blockchain server execution
package main

// Grpc tutorial
// https://www.youtube.com/watch?v=gju-bml4kdw
// Prometheus tutorial
// https://blog.alexellis.io/prometheus-monitoring/
// many helping video /Zoom/2018-10-26 13.42.32 Matias Berrueta's Personal Meeting Room ...
import (
	"fmt"
	"github.com/mberrueta/test1/metrics"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/mberrueta/test1/blockchain"
	"github.com/mberrueta/test1/proto"
	"github.com/mberrueta/test1/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	metrics.SetupMetrics()

	startGRPC()

	startHTTP()
}

func startGRPC() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	proto.RegisterBlockchainServer(s, &server.Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	grpc_prometheus.Register(s)

	go func() {
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalf("unable to listen port 8080: %v", err)
		}

		fmt.Printf("listen port: %v\n", "8080")

		s.Serve(listener)
	}()
}

func startHTTP() {
	http.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Addr:           ":9999",
		ReadTimeout:    8 * time.Second,
		WriteTimeout:   8 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := srv.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	}

	fmt.Printf("listen port: %v\n", "9999")
}
