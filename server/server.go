package server

import (
	"github.com/mberrueta/test1/blockchain"
	"github.com/mberrueta/test1/metrics"
	"github.com/mberrueta/test1/proto"
	"github.com/prometheus/client_golang/prometheus"

	"golang.org/x/net/context"
)

type Server struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock : adds new block to blockchain
func (s *Server) Add(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	labels := prometheus.Labels{"method": "Add", "op": ""}
	timer := prometheus.NewTimer(metrics.Duration.With(labels))
	defer timer.ObserveDuration()
	metrics.Executing.With(labels).Inc()
	defer metrics.Executing.With(labels).Dec()
	defer metrics.Hits.With(labels).Inc()

	block := s.Blockchain.AddBlock(in.Data, in.Verification)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain : returns blockchain
func (s *Server) List(ctx context.Context, in *proto.GetBlockRequest) (*proto.GetBlockResponse, error) {
	labels := prometheus.Labels{"method": "List", "op": ""}
	timer := prometheus.NewTimer(metrics.Duration.With(labels))
	defer timer.ObserveDuration()
	metrics.Executing.With(labels).Inc()
	defer metrics.Executing.With(labels).Dec()
	defer metrics.Hits.With(labels).Inc()

	resp := new(proto.GetBlockResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Data:          b.Data,
			Hash:          b.Hash,
		})
	}

	return resp, nil
}
