#!/bin/bash
rm blockchain
rm *.db

#go build main.go block.go blockchain.go proofofwork.go utils.go cli.go commands.go

go build -o blockchain *.go
./blockchain
