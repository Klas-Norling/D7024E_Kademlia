package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"testing"
)

func TTest_put(input []string, kademlia *Kademlia) {
	input = []string{"put", "1234klas"}
	inputHandler(input, &kademlia.rt, kademlia)
}

func TTest_get(input string, kademlia *Kademlia) {
	_input := []byte(input)
	hash := sha1.New()
	hash.Write(_input)
	sha1_data := hex.EncodeToString(hash.Sum(nil))

	inputArray := []string{"get", sha1_data}
	fmt.Println(sha1_data)

	inputHandler(inputArray, &kademlia.rt, kademlia)
}

func TTest_exit(input string, kademlia *Kademlia) {
	inputArray := []string{"exit"}
	inputHandler(inputArray, &kademlia.rt, kademlia)
}

func TestCli(t *testing.T) {
	node := InitializeNode()
	input := []string{}

	TTest_put(input, &node)
	TTest_get("1234klas", &node)

}

// Using stuff from the kademlia package here. Something like...
// node := kademlia.InitializeNode()
//fmt.Println("Hello")
//kademlia.NewKademliaID("a1b9bdfcb1f469376df7431bbb2a375fb3fb413a")

// kademlia.Cli(&node)
// kademlia.Test_put("put klasnorling", &node)
// kademlia.Test_get("klasnorling", &node)
// kademlia.Test_exit("exit k", &node)
