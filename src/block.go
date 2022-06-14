package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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
const GenesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type Block struct {
	Version       uint64
	PrevBlockHash []byte //前一个hash
	MerkleRoot    []byte //先填写为空，后续v4的时候使用
	TimeStamp     uint64 //从1970.1.1至今的秒数
	Difficulity   uint64 //挖矿的难度值, v2时使用
	Nonce         uint64 //随机数，挖矿找的就是它!
	Hash          []byte //当前区块哈希, 区块中本不存在的字段，为了方便我们添加进来
	Data          []byte //数据，目前使用字节流，v4开始使用交易代替
}

// 创建区块 对Block的每一个字段填充数据
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:       00,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    []byte{},
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulity:   10,
		Nonce:         10,
		Hash:          []byte{},
		Data:          []byte(data),
	}
	//v1版本
	// block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//为了生成区块哈希，我们实现一个简单的函数，来计算哈希值，没有随机值，没有难度值
func (block *Block) SetHash() {
	// var data []byte
	// data = append(data, block.PrevBlockHash...)
	// data = append(data, block.Data...)
	tmp := [][]byte{
		uintToByte(block.Version),
		block.PrevBlockHash,
		block.MerkleRoot,
		uintToByte(block.TimeStamp),
		uintToByte(block.Difficulity),
		block.Data,
		uintToByte(block.Nonce),
	}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

//序列化, 将区块转换成字节流
func (block *Block) Serialize() []byte {

	var buffer bytes.Buffer

	//定义编码器
	encoder := gob.NewEncoder(&buffer)

	//编码器对结构进行编码，一定要进行校验
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

func Deserialize(data []byte) *Block {

	//fmt.Printf("解码传入的数据: %x\n", data)

	var block Block

	//创建解码器
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return &block
}
