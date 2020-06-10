package BlockBuffer

import (
  "fmt"
//  "sync"
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  b1 "../Block"
)

var (
  BlkBuffer BlockBuffer
)

type BlockBuffer struct {
  Size int
  Hash []string
  Body []b1.Block
}

func (a *BlockBuffer) InsertBlock(b b1.Block) {
  a.Hash = append(a.Hash, GenerateHash(b))
  a.Body = append(a.Body, b)
  a.Size += 1
}

func (a *BlockBuffer) FindBlock(hash string) (bool, int) {
  for i := 0; i <len(a.Hash); i++ {
    if (a.Hash[i] == hash) {
      return true, i
    }
  }
  return false, -1
}

func (a *BlockBuffer) GetBlock(index int) (string, b1.Block) {
  if (index < a.Size) {
    return a.Hash[index], a.Body[index]
  } else {
    fmt.Println("BlockBuffer: Index Out Of Range")
    return "", b1.Block{}
  }
}

func (a *BlockBuffer) RemoveBlock(index int) (string, b1.Block) {
  if (index < a.Size) {
    hash := a.Hash[index]
    body := a.Body[index]

    a.Hash = append(a.Hash[:index], a.Hash[index+1:]...)
    a.Body = append(a.Body[:index], a.Body[index+1:]...)

    a.Size -= 1

    return hash, body
  } else {
    fmt.Println("BlockBuffer: Index Out Of Range")
    return "", b1.Block{}
  }
}

func GenerateHash(blk b1.Block) string {
  blk.No = 0
  blk.Miner = ""
  blk.Time = ""
  blk.Hash = ""
  blk.Prev = nil
  blk.Next = nil

  val, err := json.Marshal(blk)
  if err != nil {
    fmt.Println("Error:", err)
  }
  hash := sha256.Sum256(val)
  return hex.EncodeToString(hash[:])
}

func (a BlockBuffer) PrintBlockBuffer() {
  for i := 0; i < a.Size; i++ {
    fmt.Println("Block#", i)
    fmt.Println("Hash:", a.Hash[i])
    fmt.Println("Body:", a.Body[i])
  }
}
