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
	*/

	if returnIpAddress() == "172.16.238.10:8080" {
		root_id := kademlia.NewKademliaID(generateHashForRootNode())
		contact := kademlia.NewContact(root_id, "172.16.238.10:8080")
		rt := kademlia.NewRoutingTable(contact)
		//numberofreplicas := 0
		//kademlia.Listen("172.16.238.10", 8080, &numberofreplicas, rt)
		go kademlia.NewListenFunc("172.16.238.10:8080", rt)

	} else {
		node_id := kademlia.NewKademliaID(generateHashforNode())
		ipaddress := returnIpAddress()
		//ip, port := getIpPort(ipaddress)
		contact := kademlia.NewContact(node_id, ipaddress)
		rt := kademlia.NewRoutingTable(contact)

		go kademlia.NewListenFunc(ipaddress, rt)

		//sendingstring := ([]byte("find_node" + ";" + ipaddress))
		//fmt.Println(ipaddress)
		time.Sleep(3 * time.Second)
		contacts := rt.FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 20)
		fmt.Println("BEFORE:------------------")
		for i := range contacts {
			fmt.Println("BEFORE:------------------", contacts[i].String())
		}
		//returned_contacts := kademlia.InitiateSender("172.16.238.10:8080", sendingstring, rt)
		UNUSED(rt)
		contacts = rt.FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 20)
		for i := range contacts {
			fmt.Println("AFTER IN MAIN:", contacts[i].String())
		}
		//fmt.Println("returned contacts: ", returned_contacts)

	}

	//NODELOOKUP

	time.Sleep(10 * time.Second)
	fmt.Println("Closed down")

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

func generateHashForRootNode() string {
	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string("172.16.238.10:8080")))
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
	fmt.Println("LAST NUMBER VALUES", last_number)

	port1, err := strconv.Atoi("8080")
	port2, err := strconv.Atoi(fmt.Sprintf("%v", last_number))

	port := port1 + port2

	address := fmt.Sprintf("%v:", hostid[0]) + fmt.Sprintf("%v", port)

	UNUSED(err1, err)
	fmt.Println("IPADDRESS: ", address)
	if address == "172.16.238.10:8090" {
		return "172.16.238.10:8080"
	}

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

//kademlia.InitiateSender("172.16.238.10:8080", sendingstring, rt)
/*
	node_id := kademlia.NewKademliaID(generateHashforNode())
	ipaddress := returnIpAddress()
	//ip, port := getIpPort(ipaddress)
	contact := kademlia.NewContact(node_id, ipaddress)
	rt := kademlia.NewRoutingTable(contact)

	go kademlia.NewListenFunc(ipaddress, rt)

	sendingstring := ([]byte("find_node" + ";" + ipaddress))
*/

func nodelookup_func(dst_address string, target_address string, rt *kademlia.RoutingTable) []kademlia.Contact {
	mocking_rt := *rt
	closest_contacts := mocking_rt.FindClosestContacts(kademlia.NewKademliaID(generateHashforTargetNode(target_address)), 3)

	var shortlist = kademlia.ContactCandidates{}
	shortlist.Append(closest_contacts)

	/*
		for i := range closest_contacts {
			fmt.Println(closest_contacts[i].String())
		}*/

	var contacted_nodes_array []kademlia.Contact

	sendingString := ""
	if len(shortlist.GetContacts(20)) == 0 {
		closest_contacts = append(closest_contacts, kademlia.NewContact(kademlia.NewKademliaID(generateHashforNode()), returnIpAddress()))
	} else {
		c := make(chan []kademlia.Contact, 3)
		length_of_array := len(closest_contacts)
		k := 1
		j := 2
		for i := range closest_contacts {

			if check_if_node_exists(contacted_nodes_array, closest_contacts[i]) == true {
				fmt.Println("Already contacted")
			} else {
				if i >= length_of_array {
					break
				}

				ipaddr := closest_contacts[i].Address
				if k >= length_of_array {
					break
				}

				ipaddr2 := closest_contacts[k].Address
				if j >= length_of_array {
					break
				}
				ipaddr3 := closest_contacts[j].Address

				go kademlia.InitiateSender(ipaddr, []byte(sendingString), rt, c)
				go kademlia.InitiateSender(ipaddr2, []byte(sendingString), rt, c)
				go kademlia.InitiateSender(ipaddr3, []byte(sendingString), rt, c)
				contacted_nodes_array = append(contacted_nodes_array, closest_contacts[i])

				x, y, z := <-c, <-c, <-c

				shortlist.Append(x)
				shortlist.Append(y)
				shortlist.Append(z)
				shortlist.Sort()
				i += 3
				k += 3
				j += 3

			}

		}

	}

	return closest_contacts
}

func check_if_node_exists(contacts []kademlia.Contact, target kademlia.Contact) bool {
	for _, contact := range contacts {
		if contact == target {
			return true // Contact exists
		}
	}
	return false
}

func getIpPort(address string) (ip string, port int) {
	port_number, err := strconv.Atoi(address[len(address)-4:])
	UNUSED(err)
	ip_address := address[:len(address)-5]
	fmt.Println("port number: ", port_number, "ip_address: ", ip_address)
	return ip_address, port_number
}

func generateHashforTargetNode(target_address string) string {

	address := target_address

	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string(address)))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))

	return sha1_addrs
}
