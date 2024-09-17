// TODO: Add package documentation for `main`, like this:
// Package main something something...
package main

import (
	"crypto/sha1"
	"d7024e/kademlia"
	"encoding/hex"
	"fmt"

	// "math/big"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	time.Sleep(400 * time.Second)
	fmt.Println("Pretending to run the kademlia app...")
	// Using stuff from the kademlia package here. Something like...
	// contact := kademlia.NewContact(id, "localhost:8000")

	//bucket := kademlia.NewBucket()
	//bucket.AddContact()

	// fmt.Println(contact.String())
	// fmt.Printf("%v\n", contact)
	// generateNodes(contact)

	hostname, err := os.Hostname()
	fmt.Println("Hostname: ", hostname, "Error: ", err)

	//addrs, err := net.InterfaceAddrs()
	//fmt.Println("ip: ", addrs, "err: ", err)

	// Hashes the hostname to 160 bits (in hex)

	hash := sha1.New()
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string("172.16.238.10")))
	hash.Write([]byte(string(hostname)))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))
	fmt.Println("Hashed hostname: ", sha1_hash)
	fmt.Println("Hashed ipaddress: ", sha1_addrs)
	id := kademlia.NewKademliaID(sha1_addrs)
	contact := kademlia.NewContact(id, "172.16.238.10:8080")
	fmt.Println("Hashed ipaddress: ", contact)
	hostid, err := net.LookupIP(hostname)
	varnetlook, err := net.LookupPort("ip", "http")
	fmt.Println("what port: ", varnetlook)

	fmt.Println("ipaddress: ", hostid[0])
	time.Sleep(500 * time.Second)
}

func generateNodes(contact kademlia.Contact) {
	nodeId := generateRandomID(int64(rand.Intn(100)), 160)
	id := kademlia.NewKademliaID(nodeId)

	bucket := kademlia.NewBucket()
	bucket.AddContact(contact)
	fmt.Println("Node id ", id)
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
