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
func initializeNode() Kademlia {
	id := NewKademliaID(generateHashforNode())
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
	// TODO
}

// Takes a 160 bit key, if the corresponding value is present on the recipient,
// the associated data is returned
func (kademlia *Kademlia) LookupData(hash string) (data []byte) {
	// TODO
	value, exists := kademlia.data[hash]

	if exists {
		fmt.Println("Value already exists: ", value)

	} else {
		fmt.Println("Value does not exists, searching for value in k closest nodes")

		var closest_contacts []Contact
		k := 3

		// Finding the three closest nodes
		closest_contacts = kademlia.rt.FindClosestContacts(kademlia.rt.me.ID, k)

		for i := 0; i < len(closest_contacts); i++ {
			contact := closest_contacts[i]
			data := kademlia.network.SendFindDataMessage(hash, contact)

			kademlia.Store(data)

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
