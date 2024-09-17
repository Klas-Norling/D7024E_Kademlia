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
	time.Sleep(1 * time.Second)
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
	hostid, err := net.LookupIP(hostname)

	id := kademlia.NewKademliaID(sha1_addrs)
	contact := kademlia.NewContact(id, "172.16.238.10:8080")

	//id2 := kademlia.NewKademliaID(sha1_hash)
	ip_array := string(hostid[0])
	last_number := ip_array[len(ip_array)-1]
	fmt.Println("ip array: ", ip_array, "last number: ", last_number)
	port := "808" + fmt.Sprintf("%v", last_number)
	address := fmt.Sprintf("%v", hostid[0]) + port
	fmt.Println("port: ", port)
	fmt.Println("Address: ", address)
	//contact2 := kademlia.NewContact(id2, address)

	fmt.Println(hostid)

	fmt.Println("Hashed ipaddress: ", contact)

	fmt.Println("ipaddress: ", hostid[0])
	time.Sleep(50 * time.Second)
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
