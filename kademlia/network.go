package kademlia

type Network struct {
}

func Listen(ip string, port int) {
	// TODO
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string, contact Contact) (data []byte) {
	// TODO
	return
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
