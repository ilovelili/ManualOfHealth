package main

import (
	"context"
	"flag"
	"log"
	"time"

	proto "github.com/ilovelili/manualofhealth/proto"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "list blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("faile to dail server localhost:8888: %v", err)
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}

}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})

	if err != nil {
		log.Fatalf("unable to add block: %v", err)
	}

	log.Println("New block hash: ", block.GetHash())
}

func getBlockchain() {
	blockchain, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to get blockchain: %v", err)
	}

	for _, block := range blockchain.Blocks {
		log.Printf("Hash: %s, PrevBlockHash: %s, Data: %s", block.Hash, block.PrevBlockHash, block.Data)
	}
}
