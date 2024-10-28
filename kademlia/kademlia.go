package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
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

	//go NewListenFunc(contact.Address, &rt)
	time.Sleep(time.Second * 1)

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
	value, exists := kademlia.data[hash]
	fmt.Println("hash value: ", hash)
	fmt.Println(kademlia.data)
	new_hash := NewKademliaID(hash)

	if exists {
		fmt.Println("Value already exists: ", string(value))

	} else {
		fmt.Println("Value does not exists, searching for value in k closest contacts")

		var closest_contacts []Contact
		k := 3

		// Finding the three closest nodes
		closest_contacts = kademlia.rt.FindClosestContacts(new_hash, k)
		fmt.Println(closest_contacts)
		for i := 0; i < len(closest_contacts); i++ {
			fmt.Println(closest_contacts)
			contact := closest_contacts[i]
			data := kademlia.network.SendFindDataMessage(hash, contact)

			if Hash(data) == hash {
				fmt.Println("is the hash true:", Hash(data), " ==? ", hash)
				fmt.Println("Stored value: ", data)
				fmt.Println("value fetched from:", contact.Address)
				bytes_in_data := []byte(data)
				kademlia.Store(bytes_in_data)
				break
			}

		}

		//remove UNUSED
		UNUSED(closest_contacts)

	}

	return value
}

func Hash(data string) string {
	byte_data := []byte(data)
	hash := sha1.New()
	hash.Write(byte_data)
	sha1_data := hex.EncodeToString(hash.Sum(nil))
	return sha1_data
}

// Takes data and hashes it to a 160 bit key then converts into hexadecimal
// Note: Need to add a handle RPC file seperates lookupdata, store and lookupcontact
func (kademlia *Kademlia) Store(data []byte) {
	fmt.Println("hellostore")
	fmt.Println("data in store:", data)
	hash := sha1.New()
	hash.Write(data)
	sha1_data := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("sha1_data: ", sha1_data)
	kademlia.data[sha1_data] = data
	fmt.Println("hellostor3")

}

func GetRoutingTableMe(rt RoutingTable) Contact {
	return rt.me
}

func (kademlia *Kademlia) GetRoutingtable() RoutingTable {
	return kademlia.rt
}
