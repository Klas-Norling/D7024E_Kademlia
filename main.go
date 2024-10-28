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
	"strings"

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
		time.Sleep(time.Second * 3)

		kademlia.SendPingMessage(&root_contact, &contact)
		time.Sleep(15 * time.Second)
		Join(&kad)
		go kademlia.Cli(&kad)
		go kademlia.NewListenFunc(returnIpAddress(), &rt, &kad)

	} else {

		go kademlia.Cli(&kad)
		go kademlia.NewListenFunc(returnIpAddress(), &rt, &kad)
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
	array := fixPortInShortlist(shortlist.GetContacts(20))

	fmt.Println("Closest contacts final ound:", array)
	fmt.Println("before foorloop", rt.FindClosestContacts(kademlia.NewKademliaID(GenerateHashforNode()), 4))
	for i := range array {
		rt.AddContact(array[i])
	}
	fmt.Println("before foorloop", rt.FindClosestContacts(kademlia.NewKademliaID(GenerateHashforNode()), 4))
	return closest_contacts
}

func fixPortInShortlist(shortlist []kademlia.Contact) []kademlia.Contact {
	var root_index int
	for i := range shortlist {
		fmt.Println(shortlist[i].Address)
		split := strings.Split(shortlist[i].Address, ":")
		ipaddress := split[0]
		spliceIp := strings.Split(ipaddress, ".")
		last_number_in_ip, err := strconv.Atoi(spliceIp[3])
		new_portnumber := 8080 + last_number_in_ip
		new_ip := ipaddress + ":" + strconv.Itoa(new_portnumber)
		shortlist[i].Address = new_ip
		if ipaddress == "172.16.238.10" {
			shortlist[i].Address = "172.16.238.10:8080"
		}
		fmt.Println("returnipaddress: ", returnIpAddress())
		if returnIpAddress() == new_ip {
			fmt.Println("In returnipadress if statement")
			root_index = i
		}
		UNUSED(err)
	}
	shortlist = removeElementFromArray(shortlist, root_index)
	return shortlist
}

func removeElementFromArray(shortlist []kademlia.Contact, index int) []kademlia.Contact {
	return append(shortlist[:index], shortlist[index+1:]...)

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
