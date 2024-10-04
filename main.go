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

	fmt.Println("Pretending to run the kademlia app...")
	//generate hash and create contact for our node.
	/*	id_forOurNode := kademlia.NewKademliaID(generateHashforNode())
		contact_OurNode := kademlia.NewContact(id_forOurNode, returnIpAddress())

		//create a routing table for our node that has the root node and our node
		rt := kademlia.NewRoutingTable(contact_OurNode)

		address := returnIpAddress()

		go kademlia.NewListenFunc()*/
	if returnIpAddress() == "172.16.238.10:8090" {
		root_id := kademlia.NewKademliaID(generateHashForRootNode())
		contact := kademlia.NewContact(root_id, "172.16.238.10:8080")
		rt := kademlia.NewRoutingTable(contact)
		numberofreplicas := 0
		kademlia.Listen("172.16.238.10", 8080, &numberofreplicas, rt)
	}

	//TestRoutingTable()
	/*
		me := kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
		rt := kademlia.NewRoutingTable(me)
		rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("eb70d2b212be125aaa890c4082f44084d5a00180"), "172.16.238.10:8001"))
		addcontacts(rt)*/

	//go rt.FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 20)
	//go rt.FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 20)
	//go rt.FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 20)
	/*
		In main, create all the contacts in the system at the beginning of the run,
		later when a new node joins, make it so that all other nodes get updated.
	*/

	/*
		Varje nod ska kontakta rootnoden genom nodelookup, rootnoden ska sedan skicka till dem tillbakavad dem fick från nodelookup.

		Network ska fungera att alla noder är i standby för att lyssna på sina porter och kommer då få ett komandosom de ska köra
	*/
	/*
		if returnIpAddress() == "172.16.238.10:8090" {
			// save contacts/create contacts
			// routing table needs to store contacts

			numberofreplicas := 0

			root_id := kademlia.NewKademliaID(generateHashForRootNode())
			contact := kademlia.NewContact(root_id, "172.16.238.10:8080")
			rt := kademlia.NewRoutingTable(contact)

			//these contacts should be static
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

			kademlia.Listen("172.16.238.10", 8080, &numberofreplicas, rt)
		} else {
			time.Sleep(1 * time.Second)
			id_Root_Node := kademlia.NewKademliaID(generateHashForRootNode())

			//generate a contact to the rootnode
			contact_RootNode := kademlia.NewContact(id_Root_Node, "172.16.238.10:8080")

			//generate hash and create contact for our node.
			id_forOurNode := kademlia.NewKademliaID(generateHashforNode())
			contact_OurNode := kademlia.NewContact(id_forOurNode, returnIpAddress())

			//create a routing table for our node that has the root node and our node
			rt := kademlia.NewRoutingTable(contact_OurNode)
			rt.AddContact(contact_RootNode)

			//these contacts should be static
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
			rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

			// test(&contact_RootNode, &contact_OurNode)

			kademlia.Join("172.16.238.10:8080", rt)

		}
	*/

}

func test(contact_root *kademlia.Contact, contact_own *kademlia.Contact) {
	fmt.Println(returnIpAddress())
	if returnIpAddress() == "172.16.238.10:8090" {
		//listen

		// kademlia.Listen("172.16.238.10", 8080)
	} else {
		time.Sleep(1 * time.Second)
		kademlia.SendPingMessage(contact_root, contact_own)

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
	fmt.Println("IPADDRESS: ", address)
	return address
}

func UNUSED(x ...interface{}) {}

func TestRoutingTable() {
	rt := kademlia.NewRoutingTable(kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	//added
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("FA5E1A4DF381D0B650F5F55E8D7155719602E5A2"), "localhost:8001"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("B36828398E513AE808E0C63582FB5DBA635D7D15"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("C0932E562C38612464924C94F9114CFA3359FCAA"), "localhost:8003"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("87DEDEC92E0CEC702F31C8483F7C4B1282817CFB"), "localhost:8004"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1CFA6FA82F344CEF1269A3D746BDD56D640B209C"), "localhost:8005"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("4595501B6DD9270F9319FCC5D80F066BAA7AD885"), "localhost:8006"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("126C842B9C1548B0525DC8EC9FEA17F7813C2CB4"), "localhost:8007"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("78EA7516ED45FF89F9147494F6B3DCCE138407E9"), "localhost:8008"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("0A21410AC1C7E6C30DCF1CE7F66D479586FA7509"), "localhost:8009"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("E54E071691394B677D6A7E061ACA3A8579F05B2C"), "localhost:8010"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1745E1E0EE1EE9BEEFB44C5F75074A71C57E83A8"), "localhost:8011"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("F7537E70EDC525FA87B452F40276137DFE76D5F5"), "localhost:8012"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("7AF1EDF9CFA3EBA5929C2EAE87EB9F2FB9A008BB"), "localhost:8013"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("839C72A968674AC66D6D01F79F3DF7770AF12018"), "localhost:8014"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("6A3F114CF83CCD3E0F2E5F2DFE0C8A242B3D1A7C"), "localhost:8015"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("B8DC1D934B496E9962B150ED579165449241E6DB"), "localhost:8016"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1E7C19EB61FD4A808272FFC07090E266B2F74183"), "localhost:8017"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("78E8D1E2591845F2A6408611EA53304C4C7DA9DB"), "localhost:8018"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("B15483EC1090C84743E27CAD456A037881C79F42"), "localhost:8019"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("F10C7E4A831D9C0083371CC1077A74F4086ACC89"), "localhost:8020"))

	fmt.Println("hello")
	contacts := rt.FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}

	// TODO: This is just an example. Make more meaningful assertions.
	if len(contacts) != 6 {
		fmt.Println("Expected 6 contacts but instead got %d", len(contacts))
	}
}

func startup_routingTable(me kademlia.Contact) {

}

func addcontacts(rt *kademlia.RoutingTable) {
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	//added
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("FA5E1A4DF381D0B650F5F55E8D7155719602E5A2"), "localhost:8001"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("B36828398E513AE808E0C63582FB5DBA635D7D15"), "localhost:8002"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("C0932E562C38612464924C94F9114CFA3359FCAA"), "localhost:8003"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("87DEDEC92E0CEC702F31C8483F7C4B1282817CFB"), "localhost:8004"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1CFA6FA82F344CEF1269A3D746BDD56D640B209C"), "localhost:8005"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("4595501B6DD9270F9319FCC5D80F066BAA7AD885"), "localhost:8006"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("126C842B9C1548B0525DC8EC9FEA17F7813C2CB4"), "localhost:8007"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("78EA7516ED45FF89F9147494F6B3DCCE138407E9"), "localhost:8008"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("0A21410AC1C7E6C30DCF1CE7F66D479586FA7509"), "localhost:8009"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("E54E071691394B677D6A7E061ACA3A8579F05B2C"), "localhost:8010"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1745E1E0EE1EE9BEEFB44C5F75074A71C57E83A8"), "localhost:8011"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("F7537E70EDC525FA87B452F40276137DFE76D5F5"), "localhost:8012"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("7AF1EDF9CFA3EBA5929C2EAE87EB9F2FB9A008BB"), "localhost:8013"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("839C72A968674AC66D6D01F79F3DF7770AF12018"), "localhost:8014"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("6A3F114CF83CCD3E0F2E5F2DFE0C8A242B3D1A7C"), "localhost:8015"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("B8DC1D934B496E9962B150ED579165449241E6DB"), "localhost:8016"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("1E7C19EB61FD4A808272FFC07090E266B2F74183"), "localhost:8017"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("78E8D1E2591845F2A6408611EA53304C4C7DA9DB"), "localhost:8018"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("B15483EC1090C84743E27CAD456A037881C79F42"), "localhost:8019"))
	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("F10C7E4A831D9C0083371CC1077A74F4086ACC89"), "localhost:8020"))

}
