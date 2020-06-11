package Network

import (
	"fmt"
	"log"
	"net"
	"os"
	//	"sync"
	"bytes"
	"encoding/gob"
	"strconv"

	b1 "../Block"
	buff1 "../BlockBuffer"
	c1 "../Blockchain"
)

const (
	CommandLength = 12
	Network       = "tcp"
)

var (
	OwnAddress      string
	DefaultPeer     string
	DefaultDatabase string
	KnownNodes      = map[string]string{}
)

type Addr struct {
	AddrList    []string
	DefaultPeer string
}

type Hashes struct {
	HashList []string
}

type BlockSender struct {
	BlockList []b1.Block
}

func StartServer() {
	c1.RegisterAllGobInterfaces()

	//	c1.Chain1 = c1.LoadBlockchain()

	conn, err := net.Listen(Network, OwnAddress)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	go turnOnServer(conn)

	fmt.Println("Server has Started ...")
	AddToKnownNode(OwnAddress)
	AddToKnownNode(DefaultPeer)

	AskNodes()
}

// FUNCTIONS UNDER CONSTRUCTION BEGINS...

func HandleRegCourse(conn net.Conn, data []byte) {

}

func HandleBlock(conn net.Conn, data []byte) {

	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		log.Println(err)
	}

	var blkList BlockSender
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&blkList)
	if err != nil {
		log.Println(err)
	}

	for i, blk := range blkList.BlockList {
		if blk.Status == false {
			hash := buff1.GenerateHash1(blk)
			found, _ := buff1.BlkBuffer.FindBlock(hash)
			fmt.Println(found, i)
			if !found {
				fmt.Println(hash)
				buff1.BlkBuffer.InsertBlock(blk)

				BroadCastBlock(blk)
			}
		} else {
			hash := buff1.GenerateHash(blk)
			found := c1.Chain1.FindBlock(hash)
			fmt.Println(found)
			if !found {
				fmt.Println(hash)
				c1.Chain1.AddBlock(blk)

				BroadCastBlock(blk)
			}
		}
	}
}

func BroadCastBlock(block b1.Block) {
	for _, address := range KnownNodes {
		go SendBlock(address, block)
	}
}

func SendBlock(address string, block b1.Block) {
	blk := BlockSender{BlockList: []b1.Block{block}}
	byteBlk := GobEncode(blk)
	data := append(CmdToBytes("block"), byteBlk...)

	SendData(address, data)
}

// FUNCTIONS UNDER CONSTRUCTION ENDS...

func turnOnServer(conn net.Listener) {
	for {
		ln, err := conn.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConnection(ln)
	}
}

func handleConnection(conn net.Conn) {
	data := make([]byte, 8192)
	n, err := conn.Read(data)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	fmt.Printf("Successfully read %d bytes from client\n", n)
	cmd := BytesToCmd(data[:CommandLength])

	switch cmd {
	case "addr":
		HandleAddr(conn, data[CommandLength:])
	case "block":
		HandleBlock(conn, data[CommandLength:])
	case "askaddr":
		HandleAskAddress(conn, data[CommandLength:])
	case "rmblk":
		HandleRemoveBlock(conn, data[CommandLength:])
	case "RegCourse":
		HandleRegCourse(conn, data[CommandLength:])
	default:
		fmt.Println("Unknown Command Found")
	}
}

func AskNodes() {
	keys := []string{OwnAddress}
	addr := Addr{AddrList: keys}

	byteAddr := GobEncode(addr)

	data := append(CmdToBytes("askaddr"), byteAddr...)
	for _, v := range KnownNodes {
		go SendData(v, data)
	}
}

func HandleAskAddress(conn net.Conn, data []byte) {
	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		log.Println(err)
	}

	var address Addr
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&address)
	if err != nil {
		log.Println(err)
	}

	AddToKnownNode(address.AddrList[0])
	BroadCastNodes()
}

func HandleRemoveBlock(conn net.Conn, data []byte) {
	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		log.Println(err)
	}

	var hsh Hashes
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&hsh)
	if err != nil {
		log.Println(err)
	}

	status, index := buff1.BlkBuffer.FindBlock(hsh.HashList[0])
	if status == true {

		_, _ = buff1.BlkBuffer.RemoveBlock(index)
		//BroadCastRemoveBlockBuffer(hsh.HashList[0])
		//blk.Status = true
		//BroadCastBlock(blk)

	} else {
		fmt.Println("Block Not Found")
	}
}

func HandleAddr(conn net.Conn, data []byte) {
	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		log.Println(err)
	}

	var address Addr
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&address)
	if err != nil {
		log.Println(err)
	}

	broadcast := false
	for _, t := range address.AddrList {
		_, found := KnownNodes[t]
		if !found {
			KnownNodes[t] = t
			broadcast = true
		}
	}

	if broadcast == true {
		BroadCastNodes()
	}
}

func SendNodes(address string) {
	keys := []string{}
	for node, _ := range KnownNodes {
		keys = append(keys, node)
	}

	addr := Addr{AddrList: keys}
	addr.AddrList = append(addr.AddrList, OwnAddress)
	byteAddr := GobEncode(addr)
	newAddr := append(CmdToBytes("addr"), byteAddr...)

	SendData(address, newAddr)
}

func SendRemoveBlockBuffer(address string, hash string) {
	hsh := Hashes{HashList: []string{hash}}
	byteHashes := GobEncode(hsh)
	newHash := append(CmdToBytes("rmblk"), byteHashes...)

	SendData(address, newHash)
}

func BroadCastRemoveBlockBuffer(hash string) {
	for _, address := range KnownNodes {
		go SendRemoveBlockBuffer(address, hash)
	}
}

func SendData(address string, data []byte) {
	if address == OwnAddress {
		return
	}

	conn, err := net.Dial(Network, address)
	if err != nil {
		log.Printf("Error: %s\n", err)

		delete(KnownNodes, address)
		return
	}
	//  n, err := conn.Write(data)
	n, err := conn.Write(data)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	fmt.Printf("Written %d bytes successfully on %s\n", n, address)
}

func BroadCastNodes() {
	for address, _ := range KnownNodes {
		go SendNodes(address)
	}
}

func PrintKnownNodes() {
	k := 0
	for _, node := range KnownNodes {
		fmt.Println(strconv.Itoa(k)+".", node)
		k += 1
	}
}

func GobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Printf("Error5: %f\n", err)
	}

	return buff.Bytes()
}

func CmdToBytes(cmdString string) []byte {
	var bytes [CommandLength]byte

	for i, c := range cmdString {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func BytesToCmd(byteCmd []byte) string {
	var cmd []byte

	for _, c := range byteCmd {
		if c != 0x0 {
			cmd = append(cmd, c)
		}
	}

	return fmt.Sprintf("%s", cmd)
}

func AddToKnownNode(addr string) {
	KnownNodes[addr] = addr
}
