package kademlia

import (
	"encoding/hex"
	"testing"
)

func TestNewKademliaID(t *testing.T) {
	input := "a1b9bdfcb1f469376df7431bbb2a375fb3fb413a"
	expected := KademliaID{0xa1, 0xb9, 0xbd, 0xfc, 0xb1,
		0xf4, 0x69, 0x37, 0x6d, 0xf7, 0x43, 0x1b, 0xbb,
		0x2a, 0x37, 0x5f, 0xb3, 0xfb, 0x41, 0x3a}

	id := NewKademliaID(input)
	t.Logf("Testing NewKademliaID with input: %s, got: %x\n", input, *id)

	if id == nil {
		t.Error("Expected a KademliaID instance, got nil")
		return
	}

	if *id != expected {
		t.Errorf("Expected KademliaID %x, got %x", expected, *id)
	}
}

// Test NewKademliaID function
func TestNewKademliaIDD(t *testing.T) {
	idStr := "1111111111111111111111111111111111111111"
	kademliaID := NewKademliaID(idStr)

	decoded, _ := hex.DecodeString(idStr)
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != decoded[i] {
			t.Errorf("Expected byte %d to be %x, got %x", i, decoded[i], kademliaID[i])
		}
	}
}

// Test NewRandomKademliaID function
func TestNewRandomKademliaID(t *testing.T) {
	kademliaID1 := NewRandomKademliaID()
	kademliaID2 := NewRandomKademliaID()

	if kademliaID1.Equals(kademliaID2) {
		t.Error("Expected two different random KademliaIDs to not be equal")
	}
}

// Test Less function
func TestKademliaIDLess(t *testing.T) {
	id1 := NewKademliaID("0000000000000000000000000000000000000001")
	id2 := NewKademliaID("0000000000000000000000000000000000000002")

	if !id1.Less(id2) {
		t.Error("Expected id1 to be less than id2")
	}

	if id2.Less(id1) {
		t.Error("Expected id2 to not be less than id1")
	}
}

// Test Equals function
func TestKademliaIDEquals(t *testing.T) {
	id1 := NewKademliaID("1111111111111111111111111111111111111111")
	id2 := NewKademliaID("1111111111111111111111111111111111111111")
	id3 := NewKademliaID("2222222222222222222222222222222222222222")

	if !id1.Equals(id2) {
		t.Error("Expected id1 to be equal to id2")
	}

	if id1.Equals(id3) {
		t.Error("Expected id1 to not be equal to id3")
	}
}

// Test CalcDistance function
func TestCalcDistance(t *testing.T) {
	id1 := NewKademliaID("0000000000000000000000000000000000000001")
	id2 := NewKademliaID("0000000000000000000000000000000000000002")

	distance := id1.CalcDistance(id2)
	expectedDistance := NewKademliaID("0000000000000000000000000000000000000003") // XOR of 1 and 2

	if !distance.Equals(expectedDistance) {
		t.Errorf("Expected distance to be %s, got %s", expectedDistance.String(), distance.String())
	}
}

// Test String function
func TestKademliaIDString(t *testing.T) {
	idStr := "1111111111111111111111111111111111111111"
	kademliaID := NewKademliaID(idStr)

	if kademliaID.String() != idStr {
		t.Errorf("Expected string representation to be %s, got %s", idStr, kademliaID.String())
	}
}
