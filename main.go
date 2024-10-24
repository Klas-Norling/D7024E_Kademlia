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

	root_node_id := kademlia.NewKademliaID(generateHashForRootNode())
	root_ipaddress := "172.16.238.10:8080"
	root_contact := kademlia.NewContact(root_node_id, root_ipaddress)

	node_id := kademlia.NewKademliaID(GenerateHashforNode())
	ipaddress := returnIpAddress()

	//ip, port := getIpPort(ipaddress)
	contact := kademlia.NewContact(node_id, ipaddress)
	//rt := kademlia.NewRoutingTable(contact)

	kad := kademlia.InitializeNode()
	rt := kad.GetRoutingtable()

	if "172.16.238.10:8080" != returnIpAddress() {
		fmt.Println(returnIpAddress())

		//rt.AddContact(root_contact)
		//time.Sleep(time.Second * 3)
		time.Sleep(3 * time.Second)
		kademlia.SendPingMessage(&root_contact, &contact)
		time.Sleep(3 * time.Second)
		Join(&kad)
		go kademlia.NewListenFunc(returnIpAddress(), &rt)

	} else {

		go kademlia.Cli(&kad)
		go kademlia.NewListenFunc(returnIpAddress(), &rt)
		time.Sleep(time.Second * 3)
		closest_contacts := rt.FindClosestContacts(root_node_id, 20)
		fmt.Println("HELLOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO: ", closest_contacts)

	}

	time.Sleep(time.Second * 3)
	closest_contacts := rt.FindClosestContacts(root_node_id, 20)
	fmt.Println("HELLOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO: ", closest_contacts)

	for true {
	}

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

func test_nodelookup() {
	c1 := make(chan string)
	//numberofreplicas := 0
	//kademlia.Listen("172.16.238.10", 8080, &numberofreplicas, rt)
	root_node_id := kademlia.NewKademliaID(generateHashForRootNode())
	root_ipaddress := "172.16.238.10:8080"
	root_contact := kademlia.NewContact(root_node_id, root_ipaddress)

	node_id := kademlia.NewKademliaID(GenerateHashforNode())
	ipaddress := returnIpAddress()

	//ip, port := getIpPort(ipaddress)
	contact := kademlia.NewContact(node_id, ipaddress)
	rt := kademlia.NewRoutingTable(contact)
	rt.AddContact(root_contact)

	go kademlia.NewListenFunc(ipaddress, rt)

	//sendingstring := ([]byte("find_node" + ";" + ipaddress))
	//fmt.Println(ipaddress)
	time.Sleep(3 * time.Second)
	stringmsg := []byte("ping;p")

	fmt.Println("IPADDRRRRRRRRRRRRRRRRR", returnIpAddress())

	if returnIpAddress() != "172.16.238.10:8080" && returnIpAddress() == "172.16.238.3:8083" {
		fmt.Println("test_nodelookup, before sendpingmessage")
		time.Sleep(1 * time.Second)
		fmt.Println("HOW MANY TIMES")
		go kademlia.InitiateSenderForPong(contact.Address, stringmsg, rt, c1)
		y := <-c1
		fmt.Println(y)
		contacts := nodelookup_func("172.16.238.2:8082", rt)
		fmt.Println(contacts)

	} else if returnIpAddress() != "172.16.238.10:8080" {
		go kademlia.InitiateSenderForPong(contact.Address, stringmsg, rt, c1)
		fmt.Println("HOW MANY TIMES")
		y := <-c1
		fmt.Println(y)
	}
}

func generateHashForRootNode() string {
	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string("172.16.238.10:8080")))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))
	return sha1_addrs

}

func GenerateHashforNode() string {

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

func nodelookup_func(target_address string, rt *kademlia.RoutingTable) []kademlia.Contact {
	target := kademlia.NewContact(kademlia.NewKademliaID(generateHashforTargetNode(target_address)), target_address)

	closest_contacts := rt.FindClosestContacts(target.ID, 3)

	var shortlist = kademlia.ContactCandidates{}
	shortlist.Append(closest_contacts)

	/*
		for i := range closest_contacts {
			fmt.Println(closest_contacts[i].String())
		}*/

	var contacted_nodes_array []kademlia.Contact

	sendingString := "find_node;"

	c1 := make(chan []kademlia.Contact)
	c2 := make(chan []kademlia.Contact)
	c3 := make(chan []kademlia.Contact)
	length_of_array := len(closest_contacts)
	k := 1
	j := 2
	fmt.Println("HELLO IN NODELOOKUp12345")
	number_of_threads_okayed := 0
	for i := range closest_contacts {
		fmt.Println("HELLO IN NODELOOKUp2")
		fmt.Println("closest_contacts", closest_contacts, ", number i: ", i, " length of array:", length_of_array)
		fmt.Println("closest_contacts", closest_contacts)
		if check_if_node_exists(contacted_nodes_array, closest_contacts[i]) == true {
			fmt.Println("Already contacted")
			fmt.Println("HELLO IN NODELOOKUp3")
		} else {
			fmt.Println("HELLO IN NODELOOKUp4")
			if i < length_of_array {
				//break
				fmt.Println("HELLO IN NODELOOKUp5")
				ipaddr := closest_contacts[i].Address
				sendingString = "find_node;" + ipaddr
				fmt.Println("hellolookup123")
				go kademlia.InitiateSender(ipaddr, []byte(sendingString), rt, c1)

			}

			if k < length_of_array {
				//break
				fmt.Println("HELLO IN NODELOOKUp6")
				ipaddr2 := closest_contacts[k].Address
				sendingString = "find_node;" + ipaddr2
				go kademlia.InitiateSender(ipaddr2, []byte(sendingString), rt, c2)

			}

			if j < length_of_array {
				//break
				fmt.Println("HELLO IN NODELOOKUp7")
				ipaddr3 := closest_contacts[j].Address
				sendingString = "find_node;" + ipaddr3
				go kademlia.InitiateSender(ipaddr3, []byte(sendingString), rt, c3)
				number_of_threads_okayed += 1
			}

			//Initiatesender(destination_address, command_and_targetNode, routing_table, channel)

			contacted_nodes_array = append(contacted_nodes_array, closest_contacts[i])
			fmt.Println("nodelookup threads122222")
			//x, y, z := <-c, <-c, <-c
			x := <-c1
			fmt.Println("x---:", x)
			fmt.Println("helelelelel:")
			var y, z []kademlia.Contact
			fmt.Println("number_of_threads_okayed: ", number_of_threads_okayed)
			if k < length_of_array {
				y := <-c2
				fmt.Println("y---:", y)
				fmt.Println("HELLOOOOOO!!234")
			}

			fmt.Println("helelelelel:")
			if j < length_of_array {
				z := <-c3
				fmt.Println("y---:", z)
				fmt.Println("HELLOOOOOO!!1234")
			}
			fmt.Println("helelelelel:")

			fmt.Println("z---:", z)
			fmt.Println("nodelookup threads222222")
			fmt.Println("x---:", x, "y---:", y, "z---:", z)
			fmt.Println("IS IT FTER THIS?")

			shortlist.Append(x)
			if k < length_of_array {
				fmt.Println("HELLOOOOOO!!2348")
				shortlist.Append(y)

			}

			if j < length_of_array {
				fmt.Println("HELLOOOOOO!!2349")
				shortlist.Append(z)
			}

			fmt.Println("helelelel")

			////added
			//	rt.AddContact(kademlia.NewContact(kademlia.NewKademliaID("FA5E1A4DF381D0B650F5F55E8D7155719602E5A2"), "localhost:8001"))
			shortlist = sort_shortlist(shortlist, target)
			fmt.Println("HELLOOOOOOOOOOO")
			i += 3
			k += 3
			j += 3
			fmt.Println("helelelel123")
		}

	}
	array := shortlist.GetContacts(4)
	for i := range array {
		address, ip_port := getIpPort(array[i].Address)
		newip_address := address + string(return_last_number_of_ipadress())
		fmt.Println("ARE THE ADDRESSES WORKING?:", address+string(return_last_number_of_ipadress()))
		array[i].Address = newip_address
		UNUSED(ip_port)
	}

	fmt.Println("Closest contacts final ound:", array)

	return closest_contacts
}

func return_last_number_of_ipadress() int {
	//fetch our ip address
	hostname, err1 := os.Hostname()
	hostid, err := net.LookupIP(hostname)
	ip_array := string(hostid[0])

	last_number := ip_array[len(ip_array)-1]
	fmt.Println("LAST NUMBER VALUES", last_number)
	port1, err := strconv.Atoi("8080")
	port2, err := strconv.Atoi(fmt.Sprintf("%v", last_number))

	port := port1 + port2
	UNUSED(err1, err)
	return port
}

func sort_shortlist(shortlist kademlia.ContactCandidates, target kademlia.Contact) kademlia.ContactCandidates {
	/*rt := kademlia.NewRoutingTable(kademlia.NewContact(kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	closest_candidates := shortlist.NormalGetContacts()
	for i := range closest_candidates {
		rt.AddContact(closest_candidates[i])
	}
	shortlist.rt.FindClosestContacts(target.ID, 20)*/
	fmt.Println("HELELELELEL")
	closest_candidates := shortlist.NormalGetContacts()
	fmt.Println("HELELELELEL")

	for i := range closest_candidates {
		fmt.Println("closest_candidates: ", closest_candidates)
		fmt.Println("closest.candidates[i]", closest_candidates[i])
		closest_candidates[i].CalcDistance(target.ID)
	}
	fmt.Println("Shortlist: ", shortlist)
	shortlist.Sort()
	return shortlist
}

func check_if_node_exists(contacts []kademlia.Contact, target kademlia.Contact) bool {
	for _, contact := range contacts {
		if contact == target {

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

func Join(kad *kademlia.Kademlia) {
	rt := kad.GetRoutingtable()
	id := kademlia.NewKademliaID(GenerateHashforNode())
	contact := kademlia.NewContact(id, returnIpAddress())
	rt.SetMeRoutingTable(contact)

	//generate the root node
	root_node_id := kademlia.NewKademliaID(generateHashForRootNode())
	root_ipaddress := "172.16.238.10:8080"
	root_contact := kademlia.NewContact(root_node_id, root_ipaddress)

	rt.AddContact(root_contact)

	nodelookup_func(returnIpAddress(), &rt)

}
