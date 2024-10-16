package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

func TestGenerateHashForNode(t *testing.T) {
	address := returnIpAddress()

	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string(address)))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))

	if sha1_addrs != GenerateHashforNode() {
		t.Error("Expected a KademliaID instance, got nil")
		return

	}
}

func TestgenerateHashforNode(t *testing.T) {

}
