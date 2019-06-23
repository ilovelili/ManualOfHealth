package main

import (
	"log"
	"net"

	context "context"

	proto "github.com/ilovelili/manualofhealth/proto"
	"github.com/ilovelili/manualofhealth/server/blockchain"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to listen on 8888: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}

type Server struct {
	Blockchain *blockchain.Blockchain
}

func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.GetData())
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.GetBlocks(), &proto.Block{
			Hash:          b.Hash,
			PrevBlockHash: b.PrevBlockHash,
			Data:          b.Data,
		})
	}
	return resp, nil
}
