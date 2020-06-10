package Block

import (
  "fmt"
  "time"
  "crypto/sha256"
	"encoding/hex"
  "encoding/json"
)

type Block struct {
  Status bool
  Owner string
  No int
  Title string
  Name string
  Miner string
  Content interface{}
  Time string
  Hash string
  PrevHash string
  Prev *Block
  Next *Block
}

var UserName string
var PrivateKey string

func GenerateBlock(name string, cont interface{}) Block {
  blk := Block{Status: false, Owner: UserName, No: 0, Name: name, Miner: "", Content: cont, Time: time.Now().String(), Hash: "", PrevHash: "", Prev: nil, Next: nil}
  blk.Hash = blk.GenerateBlockHash()
  return blk
}

func (b Block) GenerateBlockHash() string {
  blk := b
  blk.No, blk.Time, blk.Hash, blk.PrevHash, blk.Prev, blk.Next = 0, "", "", "", nil, nil
  val, err := json.Marshal(blk)
  if err != nil {
    fmt.Println("Error:", err)
  }
  hash := sha256.Sum256(val)
	return hex.EncodeToString(hash[:])
}

func (a Block) PrintBlock() {
  val, err := json.Marshal(a)
  if err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Println(string(val))
}

/*

func (a Block) PrintBlock() {
	fmt.Printf("<==== Block# %d ====>\n", a.No)
	fmt.Printf("Block Hash: %s\n", a.Hash)
	fmt.Printf("Previous Block Hash: %s\n", a.PrevHash)
	fmt.Printf("Time: %s\n", a.Time)
	fmt.Println("Offered Courses:", a.OfferedCourse)
}

func (a Block) GetBlockHash() string {
	a.Prev = nil
	a.Next = nil
	a.No = 0
	a.Hash = ""

	str := fmt.Sprintf("%x", a.OfferedCourse)

	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func AddToBlockBuffer(blk Block) {
	BlockBuffer[blk.Hash] = blk
}

func PrintBlockBuffer() {
	fmt.Println("\t<=== Block Buffer ===>")
	for hash, _ := range BlockBuffer {
		fmt.Println("=>", hash)
	}
}
*/
