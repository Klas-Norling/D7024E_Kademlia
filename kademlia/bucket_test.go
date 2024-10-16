package kademlia

import (
	"testing"
)

// Mock for KademliaID for testing
func newTestKademliaID(value string) *KademliaID {
	return NewKademliaID(value)
}

// Test newBucket function
func TestNewBucket(t *testing.T) {
	b := newBucket()
	if b.list == nil {
		t.Error("Expected new bucket to have an initialized list")
	}
	if b.list.Len() != 0 {
		t.Error("Expected new bucket to be empty")
	}
}

// Test AddContact function
func TestAddContact(t *testing.T) {
	b := newBucket()
	contact1 := NewContact(newTestKademliaID("1111111111111111111111111111111111111111"), "127.0.0.1")
	contact2 := NewContact(newTestKademliaID("2222222222222222222222222222222222222222"), "127.0.0.2")

	// Add contact1
	b.AddContact(contact1)
	if b.list.Len() != 1 {
		t.Error("Expected bucket length to be 1 after adding a contact")
	}

	// Add contact2
	b.AddContact(contact2)
	if b.list.Len() != 2 {
		t.Error("Expected bucket length to be 2 after adding another contact")
	}

	// Add contact1 again and check if it moved to the front
	b.AddContact(contact1)
	firstElement := b.list.Front().Value.(Contact)
	if !firstElement.ID.Equals(contact1.ID) {
		t.Error("Expected contact1 to be moved to the front of the bucket")
	}
}

// Test GetContactAndCalcDistance function
func TestGetContactAndCalcDistance(t *testing.T) {
	b := newBucket()
	targetID := newTestKademliaID("9999999999999999999999999999999999999999")
	contact1 := NewContact(newTestKademliaID("1111111111111111111111111111111111111111"), "127.0.0.1")
	contact2 := NewContact(newTestKademliaID("2222222222222222222222222222222222222222"), "127.0.0.2")

	b.AddContact(contact1)
	b.AddContact(contact2)

	contacts := b.GetContactAndCalcDistance(targetID)
	if len(contacts) != 2 {
		t.Error("Expected 2 contacts in the returned list")
	}

	for _, contact := range contacts {
		if contact.distance == nil {
			t.Error("Expected contact distance to be calculated relative to the target ID")
		}
	}
}

// Test Len function
func TestBucketLen(t *testing.T) {
	b := newBucket()
	if b.Len() != 0 {
		t.Error("Expected bucket length to be 0 initially")
	}

	contact := NewContact(newTestKademliaID("1111111111111111111111111111111111111111"), "127.0.0.1")
	b.AddContact(contact)

	if b.Len() != 1 {
		t.Error("Expected bucket length to be 1 after adding a contact")
	}
}
