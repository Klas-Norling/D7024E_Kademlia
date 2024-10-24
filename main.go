// TODO: Add package documentation for `main`, like this:
// Package main something something...
package main

import (
	"crypto/sha1"
	"d7024e/kademlia"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("Pretending to run the kademlia app...")
	// Using stuff from the kademlia package here. Something like...

	// node := kademlia.InitializeNode()
	// kademlia.Join(node)
	// kademlia.Cli(&node)
	// kademlia.Test_put("put klasnorling", &node)
	// kademlia.Test_get("klasnorling", &node)
	// kademlia.Test_exit("exit k", &node)

	test_nodelookup()
}

func test_findvalue() {
	node := kademlia.InitializeNode()
	kademlia.Join(node)

}

func test_nodelookup() {
	node := kademlia.InitializeNode()
	root_node := kademlia.NewKademliaID(kademlia.GenerateHashForRootNode())
	root_contact := kademlia.NewContact(root_node, "172.16.238.10:8080")
	rt := node.GetRoutingtable()
	rt.AddContact(root_contact)
	me := kademlia.GetRoutingTableMe(rt)

	if me.Address != "172.16.238.10:8080" && me.Address == "172.16.238.3:8083" {
		fmt.Println("test_nodelookup, before sendpingmessage")
		kademlia.SendPingMessage(&root_contact, &me)
		nodelookup_func("172.16.238.2:8082", &rt)

	} else if me.Address != "172.16.238.10:8080" {
		kademlia.SendPingMessage(&root_contact, &me)
	}
}

func test_store() {

}

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

	return closest_contacts
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

func generateHashforTargetNode(target_address string) string {

	address := target_address

	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string(address)))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))

	return sha1_addrs
}
