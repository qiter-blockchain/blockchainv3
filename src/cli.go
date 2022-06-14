package main

import (
	"fmt"
	"os"
)

//使用命令行分析
//
//1. 所有的支配动作交给命令行来做
//2. 主函数只需要调用命令行结构即可
//3. 根据输入的不同命令，命令行做相应动作
//1. addBlock
//2. printChain
//
//
//
//CLI : command line的缩写
//
//type CLI struct {
//	 bc *BlockChain
//}
//
//
//

const Usage = `
	./blockchain addBlock "xxx"
	./blockchain printChain
`

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) Run() {
	cmds := os.Args
	if len(cmds) < 2 {
		fmt.Printf(Usage)
		os.Exit(1)
	}
	switch cmds[1] {
	case "addBlock":
		if len(cmds) != 3 {
			fmt.Println(Usage)
			os.Exit(1)
		}
		fmt.Printf("添加区块命令被调用，数据：%s\n", cmds[2])

		data := cmds[2]
		cli.bc.AddBlock(data)
	case "printChain":
		fmt.Println("打印区块链命令被调用")
		cli.PrintChain()
	default:
		fmt.Println("无效的命令，请检查！")
		fmt.Println(Usage)
	}
}
