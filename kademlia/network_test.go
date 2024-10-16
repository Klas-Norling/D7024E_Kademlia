package kademlia

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// Test MakeSenseOfStringMessage function
func TestMakeSenseOfStringMessage(t *testing.T) {
	message := "HELLO;WORLD"
	part1, part2 := MakeSenseOfStringMessage(message)

	if part1 != "HELLO" || part2 != "WORLD" {
		t.Errorf("Expected 'HELLO' and 'WORLD', got %s and %s", part1, part2)
	}

	singleMessage := "SINGLE"
	part1, part2 = MakeSenseOfStringMessage(singleMessage)
	if part1 != "SINGLE" || part2 != "" {
		t.Errorf("Expected 'SINGLE' and '', got %s and %s", part1, part2)
	}
}

// Test encodeContactsToBytes and decodeContactsFromBytes
func TestEncodeDecodeContacts(t *testing.T) {
	contacts := []Contact{
		NewContact(newTestKademliaID("1111111111111111111111111111111111111111"), "127.0.0.1"),
		NewContact(newTestKademliaID("2222222222222222222222222222222222222222"), "127.0.0.2"),
	}

	encodedBytes, err := encodeContactsToBytes(contacts)
	if err != nil {
		t.Errorf("Failed to encode contacts: %v", err)
	}

	decodedContacts, err := decodeContactsFromBytes(encodedBytes)
	if err != nil {
		t.Errorf("Failed to decode contacts: %v", err)
	}

	if len(decodedContacts) != len(contacts) {
		t.Errorf("Expected %d contacts, got %d", len(contacts), len(decodedContacts))
	}
}

// Test EncodeToBytes and DecodeToPerson
func TestEncodeDecodePerson(t *testing.T) {
	contacts := []Contact{
		NewContact(newTestKademliaID("1111111111111111111111111111111111111111"), "127.0.0.1"),
		NewContact(newTestKademliaID("2222222222222222222222222222222222222222"), "127.0.0.2"),
	}

	encodedBytes := EncodeToBytes(contacts)
	decodedContacts := DecodeToPerson(encodedBytes)

	if len(decodedContacts) != len(contacts) {
		t.Errorf("Expected %d contacts, got %d", len(contacts), len(decodedContacts))
	}
}

func TestSwitch_case_find_value(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))
	ipaddr := "172.16.238.10:8080"
	data := switch_case_find_value(ipaddr, ipaddr, rt)
	fmt.Println("FINDERS", data)
}
func TestSwitch_case_find_node(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))
	ipaddr := "172.16.238.10:8080"
	newip_forsender, newport_forsender := getdecrementIpPort("172.16.238.10:8080")
	contacts, contact := switch_case_find_node(newip_forsender, ipaddr, newport_forsender, rt)
	fmt.Println(contact, contacts)

}

// Mock handler for NewListenFunc test
func mockHandler(conn net.Conn, rt *RoutingTable) {
	defer conn.Close()
	// Simulate handling a request
}

// Test NewListenFunc
func TestNewListenFunc(t *testing.T) {
	// Setup mock routing table
	rt := &RoutingTable{}

	go func() {
		// Run the NewListenFunc in a goroutine to simulate a listener
		NewListenFunc("localhost:9000", rt)
	}()

	time.Sleep(1 * time.Second) // Give the listener time to start

	// Simulate a client connecting to the listener
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		t.Errorf("Failed to connect to listener: %v", err)
	}
	defer conn.Close()

	// Test connection
	if conn == nil {
		t.Error("Expected valid connection, got nil")
	}
}
