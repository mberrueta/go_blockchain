package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block keeps block headers
type Block struct {
	Data          string
	PrevBlockHash string
	Hash          string
	Verification  string
}

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates and returns Block
func NewBlock(data string, verification string, prevBlockHash string) *Block {
	sha := sha256.Sum256([]byte(prevBlockHash + data + verification))
	hash := hex.EncodeToString(sha[:])

	block := &Block{data, prevBlockHash, hash, verification}

	return block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string, verification string) *Block {
	var newBlock *Block

	if len(bc.Blocks) > 0 {
		prevBlock := bc.Blocks[len(bc.Blocks)-1]
		newBlock = NewBlock(data, verification, prevBlock.Hash)
	} else {
		newBlock = NewBlock(data, verification, "")
	}
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{}}
}
