package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Test_put(input string, kademlia *Kademlia) {
	InputHandler(input, kademlia)
}

func Test_get(input string, kademlia *Kademlia) {
	_input := []byte(input)
	hash := sha1.New()
	hash.Write(_input)
	sha1_data := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(sha1_data)
	InputHandler("get "+sha1_data, kademlia)
}

func Test_exit(input string, kademlia *Kademlia) {
	InputHandler(input, kademlia)
}
