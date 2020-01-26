package assignment01IBC_i160140

import ("fmt"
  "crypto/sha256"
  "encoding/hex"
  "strconv"
)

type Transaction struct {
  Trans []string
}

func (a *Transaction) String() string {
  str1 := ""
  for i:=0;i<len(a.Trans);i++ {
    str1 += a.Trans[i]
  }
  return str1
}

func (a *Transaction) Input() {
  str1 := ""
  fmt.Scanln(&str1)
  a.Trans = append(a.Trans, str1)
}

func (a *Transaction) Add(str1 string) {
  a.Trans = append(a.Trans, str1)
}

//Block - Block structure
type Block struct {
  Blockno int
  Trans Transaction
  PrevHash string
  PrevPointer *Block
}

func (a Block) GenerateHash() string {
  hashstr := strconv.Itoa(a.Blockno) + a.Trans.String() + a.PrevHash
  result := sha256.Sum256([]byte(hashstr))
  return hex.EncodeToString(result[:])
}

//InsertBlock - Used to insert Block into the Blockchain
func InsertBlock(trans Transaction, chainHead *Block) *Block {
  var newBlock Block

  if chainHead == nil {
    newBlock.Blockno = 0
    newBlock.PrevHash = "0000000000000000000000000000000000000000000000000000000000000000"
    newBlock.PrevPointer = nil
    newBlock.Trans = trans
  } else {
    newBlock.Blockno = chainHead.Blockno + 1
    newBlock.PrevHash = chainHead.GenerateHash()
    newBlock.PrevPointer = chainHead
    newBlock.Trans = trans
  }

  return &newBlock
}

//ListBlocks - list all the blocks in the blockchain
func ListBlocks(chainHead *Block) {
  nodePtr := chainHead

  fmt.Println("\t\tLISTING BLOCKCHAIN: ")

  for nodePtr != nil{
    fmt.Println("Block#", nodePtr.Blockno)
    fmt.Println("Previous Block Hash #", nodePtr.PrevHash)
    fmt.Println("Previous Pointer: ", nodePtr.PrevPointer)
    for i:=0;i<len(nodePtr.Trans.Trans);i++ {
      fmt.Println("Transaction# ",i," ",nodePtr.Trans.Trans[i])
    }
    nodePtr = nodePtr.PrevPointer
  }
}
