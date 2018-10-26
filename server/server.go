package server

import (
	"github.com/mberrueta/test1/blockchain"
	"github.com/mberrueta/test1/proto"

	"golang.org/x/net/context"
)

type Server struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock : adds new block to blockchain
func (s *Server) Add(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data, in.Verification)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain : returns blockchain
func (s *Server) List(ctx context.Context, in *proto.GetBlockRequest) (*proto.GetBlockResponse, error) {
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
