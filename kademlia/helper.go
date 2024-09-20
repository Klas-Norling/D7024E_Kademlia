package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
)

func generateHashforNode(ipaddress string) string {

	address := ipaddress

	//hash our ip address
	hashed_addrs := sha1.New()
	hashed_addrs.Write([]byte(string(address)))
	sha1_addrs := hex.EncodeToString(hashed_addrs.Sum(nil))

	return sha1_addrs
}
