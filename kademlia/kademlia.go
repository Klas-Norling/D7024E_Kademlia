package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Kademlia struct {
	data map[string][]byte
	rt   RoutingTable
	//Network interface?
	network Network
}

// Initializes a node, creates a routingtable and assigns struct values
func InitializeNode() Kademlia {
	id := NewKademliaID(GenerateHashforNode())
	contact := NewContact(id, returnIpAddress())

	var rt RoutingTable = *NewRoutingTable(contact)

	kademlia := Kademlia{
		//NetworkInterface: network,
		rt:   rt,
		data: make(map[string][]byte),
	}

	return kademlia

}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	contacts := kademlia.rt.FindClosestContacts(target.ID, 3)

	var shortlist = ContactCandidates{}
	var contacted_nodes []Contact

	channel := make(chan []Contact, 3)

	if len(contacts) == 0 {
		fmt.Println("No contacts in the routingtable")
		return
	}

	shortlist.Append(contacts)

	for i := range contacts {
		if contacted_node(contacted_nodes, contacts[i]) == true {
			fmt.Println("Already contacted")

		} else {
			go kademlia.network.SendFindContactMessage(contacts[i])

			contacted_nodes = append(contacted_nodes, contacts[i])
			shortlist.Append(contacts)

			sort_shortlist(shortlist, target)

		}

	}
}

func sort_shortlist(shortlist ContactCandidates, target *Contact) ContactCandidates {
	for i := range shortlist.contacts {
		shortlist.contacts[i].CalcDistance(target.ID)
	}
	return shortlist
}

func contacted_node(contacted_nodes []Contact, node Contact) bool {
	for _, contact := range contacted_nodes {
		if contact == node {
			return true // Contact exists
		}
	}
	return false
}

// Takes a 160 bit key, if the corresponding value is present on the recipient,
// the associated data is returned
func (kademlia *Kademlia) LookupData(hash string) (data []byte) {
	value, exists := kademlia.data[hash]

	if exists {
		fmt.Println("Value already exists: ", value)

	} else {
		fmt.Println("Value does not exists, searching for value in k closest contacts")

		var closest_contacts []Contact
		k := 3

		// Finding the three closest nodes
		closest_contacts = kademlia.rt.FindClosestContacts(kademlia.rt.me.ID, k)

		for i := 0; i < len(closest_contacts); i++ {
			contact := closest_contacts[i]
			data := kademlia.network.SendFindDataMessage(hash, contact)

			kademlia.Store(EncodeToBytes(data))

			fmt.Println("Stored value: ", string(data))

		}

	}

	return value
}

// Takes data and hashes it to a 160 bit key then converts into hexadecimal
// Note: Need to add a handle RPC file seperates lookupdata, store and lookupcontact
func (kademlia *Kademlia) Store(data []byte) {
	hash := sha1.New()
	hash.Write(data)
	sha1_data := hex.EncodeToString(hash.Sum(nil))

	kademlia.data[sha1_data] = data

}
