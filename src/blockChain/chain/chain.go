package chain

import "blockChain/block"

type Chain struct {
	Blocks []*block.Block
}

func (c *Chain) AddBlock(data string) {
	prevBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}

func NewChain() *Chain {
	return &Chain{
		[]*block.Block{block.NewGenesisBlock()},
	}
}
