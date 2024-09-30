package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
)

type Kademlia struct {
	data map[string][]byte
	rt   RoutingTable
	//Network interface?
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

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

// Takes data and hashes it to a 128 bit value then converts into hexadecimal
// Note: Need to add a handle RPC file seperates lookupdata, store and lookupcontact
func (kademlia *Kademlia) Store(data []byte) {
	hash := sha1.New()
	hash.Write(data)
	sha1_data := hex.EncodeToString(hash.Sum(nil))

	kademlia.data[sha1_data] = data

}
