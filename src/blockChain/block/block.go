package block

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
)

/*一个区块，Timestamp, PrevBlockHash, Hash 是区块头（block header）规范中是一个单独的数据结构*/
type Block struct {
	Timestamp     int64  //区块创建时间
	PrevBlockHash []byte //前一个区块的哈希
	Hash          []byte //当前区块的哈希
	Data          []byte //区块存储的有效信息
}

/*将Block中的部分字段相互连接起来，然后计算出一个SHA-256的哈希值*/
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

/*用于简化创建一个区块的函数*/
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(),
		prevBlockHash,
		[]byte{},
		[]byte(data),
	}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
