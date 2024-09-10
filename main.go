// TODO: Add package documentation for `main`, like this:
// Package main something something...
package main

import (
	"d7024e/kademlia"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Pretending to run the kademlia app...")
	// Using stuff from the kademlia package here. Something like...
	id := kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	contact := kademlia.NewContact(id, "localhost:8000")
	fmt.Println(contact.String())
	fmt.Printf("%v\n", contact)
	generateNodes(contact)
}

func generateNodes(contact kademlia.Contact) {
	nodeId := generateRandomID(int64(rand.Intn(100)), 160)
	id := kademlia.NewKademliaID(nodeId)

	bucket := kademlia.NewBucket()
	bucket.AddContact(contact)
	fmt.Println("Node id ")
	fmt.Printf("%v\n", nodeId)
}

func generateRandomID(seed int64, binLength int) string {
	id := ""
	rand.Seed(time.Now().UnixNano() - seed)
	for i := 0; i < binLength; i++ {
		id += strconv.Itoa(rand.Intn(2))
	}

	return id
}

func sendMessage()
