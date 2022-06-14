package main

import (
	"fmt"
	"time"
)

// 定义结构
// 前区块哈希
// 当前区块哈希
// 数据

// 创建区块
// 生成哈希
// 引入区块链
// 添加区块
// 重构代码

func main() {
	// fmt.Printf("helloworld\n")
	// block := NewBlock(genesisInfo, []byte{0x0000000000000})

	bc := NewBlockChain()
	bc.AddBlock("班主任来了，大家欢迎~")
	bc.AddBlock("班主任走了")
	bc.AddBlock("我来了，大家欢迎")
	bc.AddBlock("我走了")

	for i, block := range bc.Blocks {
		// fmt.Printf("**************%d******************\n", i)
		// fmt.Printf("PrevBlockHash = %x\n", block.PrevBlockHash)
		// fmt.Printf("Hash = %x\n", block.Hash)
		// fmt.Printf("data = %s\n", block.Data)

		fmt.Printf("+++++++++++++++ %d ++++++++++++++\n", i)
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
		fmt.Printf("MerKleRoot : %x\n", block.MerkleRoot)

		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("TimeStamp : %s\n", timeFormat)

		fmt.Printf("Difficulity : %d\n", block.Difficulity)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", block.Data)

		pow := NewProofOfWork(block)
		fmt.Printf("IsValid: %v\n", pow.IsValid())
	}

}
