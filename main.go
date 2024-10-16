// TODO: Add package documentation for `main`, like this:
// Package main something something...
package main

import (
	"d7024e/kademlia"
	"fmt"
)

func main() {
	fmt.Println("Pretending to run the kademlia app...")
	// Using stuff from the kademlia package here. Something like...
	// node := kademlia.InitializeNode()
	fmt.Println("Hello")
	kademlia.NewKademliaID("1234567890abcdef1234567890abcdef12345678")

	// kademlia.Cli(&node)
	// kademlia.Test_put("put klasnorling", &node)
	// kademlia.Test_get("klasnorling", &node)
	// kademlia.Test_exit("exit k", &node)
}
