// Package server handles the Server implementation
package server

import (
	"github.com/mberrueta/test1/blockchain"
	"github.com/mberrueta/test1/metrics"
	"github.com/mberrueta/test1/proto"

	"golang.org/x/net/context"
)

type Server struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock : adds new block to blockchain
func (s *Server) Add(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {

	i, err := metrics.Gauge("Add", func() (interface{}, error) {
		return s.Blockchain.AddBlock(in.Data, in.Verification), nil
	})

	// resp := new(proto.AddBlockResponse)
	resp := &proto.AddBlockResponse{
		Hash: i.(*blockchain.Block).Hash,
	}
	return resp, err
}

// GetBlockchain : returns blockchain
func (s *Server) List(ctx context.Context, in *proto.GetBlockRequest) (*proto.GetBlockResponse, error) {
	i, err := metrics.Gauge("List", func() (interface{}, error) {
		resp := new(proto.GetBlockResponse)

		for _, b := range s.Blockchain.Blocks {
			resp.Blocks = append(resp.Blocks, &proto.Block{
				PrevBlockHash: b.PrevBlockHash,
				Data:          b.Data,
				Hash:          b.Hash,
			})
		}

		return resp, nil
	})
	return i.(*proto.GetBlockResponse), err
}
