package kademlia

import (
	"fmt"
	"testing"
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
		NewContact(NewKademliaID("1111111111111111111111111111111111111111"), "127.0.0.1"),
		NewContact(NewKademliaID("2222222222222222222222222222222222222222"), "127.0.0.2"),
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
		NewContact(NewKademliaID("1111111111111111111111111111111111111111"), "127.0.0.1"),
		NewContact(NewKademliaID("2222222222222222222222222222222222222222"), "127.0.0.2"),
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

func TestGetIpPort(t *testing.T) {
	tests := []struct {
		address      string
		expectedIP   string
		expectedPort int
	}{
		{"192.168.0.10123", "192.168.0.10", 123},
		{"127.0.0.11234", "127.0.0.1", 1234},
		{"10.0.0.11000", "10.0.0.1", 1000},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("address=%s", test.address), func(t *testing.T) {
			ip, port := getIpPort(test.address)
			if ip != test.expectedIP || port != test.expectedPort {
				t.Errorf("getIpPort(%s) = (%s, %d); want (%s, %d)", test.address, ip, port, test.expectedIP, test.expectedPort)
			}
		})
	}
}

func TestGetDecrementIpPort(t *testing.T) {
	tests := []struct {
		address      string
		expectedIP   string
		expectedPort int
	}{
		{"192.168.0.10123", "192.168.0.10", 122},
		{"127.0.0.11234", "127.0.0.1", 1233},
		{"10.0.0.11000", "10.0.0.1", 999},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("address=%s", test.address), func(t *testing.T) {
			ip, port := getdecrementIpPort(test.address)
			if ip != test.expectedIP || port != test.expectedPort {
				t.Errorf("getdecrementIpPort(%s) = (%s, %d); want (%s, %d)", test.address, ip, port, test.expectedIP, test.expectedPort)
			}
		})
	}
}

func TestGetNewIpPort(t *testing.T) {
	tests := []struct {
		address      string
		expectedIP   string
		expectedPort int
	}{
		{"192.168.0.10123", "192.168.0.10", 124},
		{"127.0.0.11234", "127.0.0.1", 1235},
		{"10.0.0.11000", "10.0.0.1", 1001},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("address=%s", test.address), func(t *testing.T) {
			ip, port := getNewIpPort(test.address)
			if ip != test.expectedIP || port != test.expectedPort {
				t.Errorf("getNewIpPort(%s) = (%s, %d); want (%s, %d)", test.address, ip, port, test.expectedIP, test.expectedPort)
			}
		})
	}
}

func TestGetNewNEWIpPort(t *testing.T) {
	tests := []struct {
		address      string
		expectedIP   string
		expectedPort int
	}{
		{"192.168.0.10123", "192.168.0.10", 125},
		{"127.0.0.11234", "127.0.0.1", 1236},
		{"10.0.0.11000", "10.0.0.1", 1002},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("address=%s", test.address), func(t *testing.T) {
			ip, port := getNewNEWIpPort(test.address)
			if ip != test.expectedIP || port != test.expectedPort {
				t.Errorf("getNewNEWIpPort(%s) = (%s, %d); want (%s, %d)", test.address, ip, port, test.expectedIP, test.expectedPort)
			}
		})
	}
}
