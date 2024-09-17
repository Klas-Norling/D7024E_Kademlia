// TODO: Add package documentation for `main`, like this:
// Package main something something...
package main

import (
	"crypto/sha1"
	"d7024e/kademlia"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"

	// "math/big"

	"time"
)

func main() {
	//generate hash for root node
	time.Sleep(1 * time.Second)
	fmt.Println("Pretending to run the kademlia app...")
	id_Root_Node := kademlia.NewKademliaID(generateHashForRootNode())

	//generate a contact to the rootnode
	contact_RootNode := kademlia.NewContact(id_Root_Node, "172.16.238.10:8080")

	//generate hash and create contact for our node.
	id_forOurNode := kademlia.NewKademliaID(generateHashforNode())
	contact_OurNode := kademlia.NewContact(id_forOurNode, returnIpAddress())

	//create a routing table for our node that has the root node and our node
	rt := kademlia.NewRoutingTable(contact_OurNode)
	rt.AddContact(contact_RootNode)

	test(&contact_RootNode)

}

func test(contact *kademlia.Contact) {
	fmt.Println(returnIpAddress())
	if returnIpAddress() == "172.16.238.10:8090" {
		//listen

		kademlia.Listen("172.16.238.10", 8080)
	} else {
		time.Sleep(1 * time.Second)
		kademlia.SendPingMessage(contact)

	}

}

/*
func oldmain() {

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

	port1, err := strconv.Atoi("8080")
	port2, err := strconv.Atoi(fmt.Sprintf("%v", last_number))

	port := port1 + port2

	address := fmt.Sprintf("%v:", hostid[0]) + fmt.Sprintf("%v", port)
	fmt.Println("port: ", port)
	fmt.Println("Address: ", address)
	//contact2 := kademlia.NewContact(id2, address)

	fmt.Println(hostid)

	fmt.Println("Hashed ipaddress: ", contact)

	fmt.Println("ipaddress: ", hostid[0])
	time.Sleep(50 * time.Second)

}*/

func generateHashForRootNode() string {
	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string("172.16.238.10")))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))
	return sha1_addrs

}

func generateHashforNode() string {

	address := returnIpAddress()

	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string(address)))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))

	return sha1_addrs
}

func returnIpAddress() string {
	//fetch our ip address
	hostname, err1 := os.Hostname()
	hostid, err := net.LookupIP(hostname)
	ip_array := string(hostid[0])
	last_number := ip_array[len(ip_array)-1]

	port1, err := strconv.Atoi("8080")
	port2, err := strconv.Atoi(fmt.Sprintf("%v", last_number))

	port := port1 + port2

	address := fmt.Sprintf("%v:", hostid[0]) + fmt.Sprintf("%v", port)

	UNUSED(err1, err)

	return address
}

func UNUSED(x ...interface{}) {}
