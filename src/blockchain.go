package main

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const blockChainName = "blockChain.db"
const blockBucketName = "blockBucket"
const lastHashKey = "lastHashKey"

//创建区块链，使用Block数组模拟
type BlockChain struct {
	db   *bolt.DB //数据库句柄
	tail []byte   //最后一个区块的哈希值
	// Blocks []*Block
}

//实现创建区块链的方法
func NewBlockChain() *BlockChain {
	//功能分析：
	//1. 获得数据库的句柄，打开数据库，读写数据
	db, err := bolt.Open(blockChainName, 0600, nil)
	//向数据库中写入数据
	//从数据库中读取数据
	if err != nil {
		log.Panic(err)
	}
	//主函数中
	// defer db.Close()

	var tail []byte

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucketName))
		if b == nil {
			//如果b1为空，说明名字为"buckeName1"这个桶不存在，我们需要创建之
			fmt.Printf("bucket不存在，准备创建！\n")
			b, err = tx.CreateBucket([]byte(blockBucketName))

			if err != nil {
				log.Panic(err)
			}
			genesisBlock := NewBlock(GenesisInfo, []byte{})
			b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			b.Put([]byte(lastHashKey), genesisBlock.Hash)
			tail = genesisBlock.Hash
		} else {
			tail = b.Get([]byte(lastHashKey))
		}
		return nil
	})
	//创建的时候添加一个区块： 创世块

	// bc := BlockChain{Blocks: []*Block{genesisBlock}}
	return &BlockChain{db, tail}
}

// 添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 1.创建一个区块
	//bc.Blocks的最后一个区块的Hash值就是当前新区块的PrevBlockHash
	bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucketName))

		if b == nil {
			fmt.Println("bucket不存在，请检查！")
			os.Exit(1)
		}

		block := NewBlock(data, bc.tail)
		b.Put(block.Hash, block.Serialize()) //将区块序列化，转成字节流
		b.Put([]byte(lastHashKey), block.Hash)
		bc.tail = block.Hash
		return nil
	})

	// lastBlock := bc.Blocks[len(bc.Blocks)-1]
	// prevHash := lastBlock.Hash

	// 2.添加到bc.Blocks数组中
	// bc.Blocks = append(bc.Blocks, block)
}

//定义一个区块链的迭代器，包含db，current
type BlockChainIterator struct {
	db      *bolt.DB
	current []byte //当前所指向区块的哈希值
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{bc.db, bc.tail}
}

func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucketName))
		if b == nil {
			fmt.Println("bucket不存在，请检查！")
			os.Exit(1)
		}
		blockInfo := b.Get(it.current)
		block = *Deserialize(blockInfo)
		it.current = block.PrevBlockHash
		return nil
	})
	return &block
}
