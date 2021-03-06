// Package main (client) provide a client to call the blockchain
package main

import (
	"flag"
	"fmt"
	"github.com/mberrueta/test1/proto"
	"log"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.String("add", "data,code", "Add new block")
	listFlag := flag.Bool("list", false, "List all blocks")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = proto.NewBlockchainClient(conn)

	if addFlag != nil {
		fmt.Printf("addflag has the value: '%v'", *addFlag)

		addBlock(*addFlag)
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock(flag string) {
	log.Println(flag)
	split := strings.Split(flag, ",")

	block, addErr := client.Add(context.Background(), &proto.AddBlockRequest{
		Data:         split[0],
		Verification: split[1],
	})
	if addErr != nil {
		log.Fatalf("unable to add block: %v", addErr)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockchain() {
	blockchain, getErr := client.List(context.Background(), &proto.GetBlockRequest{})
	if getErr != nil {
		log.Fatalf("unable to get blockchain: %v", getErr)
	}

	log.Println("blocks:")
	for _, b := range blockchain.Blocks {
		log.Printf("hash %s, prev hash: %s, data: %s\n", b.Hash, b.PrevBlockHash, b.Data)
	}
}
