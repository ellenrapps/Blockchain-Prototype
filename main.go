package main

import (
    "fmt" 
    "bytes"         
    "crypto/sha256" 
)


type BlockInfo struct { 
    PreviousBlockHash []byte 
    CurrentBlockHash  []byte 
    Message           []byte 
}

type Blockchain struct {
    Blocks []*BlockInfo 
}


func (blockInfo *BlockInfo) CalculateBlockHash() {                                  
    headers := bytes.Join([][]byte{blockInfo.PreviousBlockHash, blockInfo.Message}, []byte{}) 
    hash := sha256.Sum256(headers)                                                               
    blockInfo.CurrentBlockHash = hash[:]  
}


func NewBlock(data string, prevBlockHash []byte) *BlockInfo {
    blockInfo := &BlockInfo{prevBlockHash, []byte{}, []byte(data)} 
    blockInfo.CalculateBlockHash()                                                           
    return blockInfo                                                              
}


func GenesisBlock() *BlockInfo {
    return NewBlock("Genesis Block", []byte{}) 
}


func (blockchain *Blockchain) AddBlock(data string) {
    previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] 
    newBlock := NewBlock(data, previousBlock.CurrentBlockHash)        
    blockchain.Blocks = append(blockchain.Blocks, newBlock)      
}


func NewBlockchain() *Blockchain { 
    return &Blockchain{[]*BlockInfo{GenesisBlock()}} 
}


func main() {
    newblockchain := NewBlockchain()    
    newblockchain.AddBlock("Transaction 1 After Genesis")  
    newblockchain.AddBlock("Transaction 2 After Genesis")
    newblockchain.AddBlock("Transaction 3 After Genesis") 
    for i, blockInfo := range newblockchain.Blocks {
		fmt.Printf("Block ID : %d \n", i)
		fmt.Printf("Message : %s\n", blockInfo.Message)
		fmt.Printf("Previous Block Hash: %x\n", blockInfo.PreviousBlockHash)                 
        fmt.Printf("Block Hash : %x\n", blockInfo.CurrentBlockHash)                         
    } 
}
