package kademlia

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Network struct {
}

func encodeContactsToBytes(contacts []Contact) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(contacts)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeContactsFromBytes(data []byte) ([]Contact, error) {
	var contacts []Contact
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&contacts)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func NewListenFunc(ip string, rt *RoutingTable) {

	ip, port := getIpPort(ip)

	ln, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))

	if err != nil {
		fmt.Println("Caught error: ", err)
		return
	}

	defer ln.Close()

	// Accept incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("Error:", err)
		}
		fmt.Println("HELLO")
		go RPC_handler(conn, rt)

		fmt.Println("Listening on ip and port", ip, port)
		time.Sleep(3 * time.Second)
		UNUSED(conn)
	}

}

func RPC_handler(conn net.Conn, rt *RoutingTable) {

	tmp := make([]byte, 1024)
	n, err := conn.Read(tmp)
	root_id := NewKademliaID(generateHashForRootNode())
	contact := NewContact(root_id, "172.16.238.10:8080")
	var kadem = Kademlia{}
	//var network_struct = Network{}

	UNUSED(err)

	receivedString := string(tmp[:n])

	switch receivedString {
	case "store":
		// Initialize or reset the store

	case "put":
		// Put key-value pair in the store

	case "get":
		// Get value by key
		address := conn.RemoteAddr().String()
		id := NewKademliaID(generateHashforNode(address))
		contact := NewContact(id, address)

		rt.AddContact(contact)
		//network_struct.SendFindContactMessage(&contact) //Useless???
		contacts := rt.FindClosestContacts(contact.ID, 20)
		bytesof_contacts, err := encodeContactsToBytes(contacts)
		NewSenderFunc(&contact, &rt.me, bytesof_contacts)
		UNUSED(err)

	case "nodelookup":
		kadem.LookupContact(&contact)

	case "valuelookup":
		//kadem.LookupData()

	default:
		return
	}

}

func NewSenderFunc(contact_other *Contact, contact_own *Contact, data []byte) {
	ip, port := getIpPort(contact_own.Address)
	fmt.Println(ip, port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}

	conn, err := dialer.Dial("tcp", contact_other.Address)
	if err != nil {
		fmt.Println("Error caught: ", err)
		defer conn.Close()
	}

	fmt.Println("Connection was established to: ", conn.RemoteAddr())

	conn.Write(data)
}

func NewSenderFunc_Nodelookup(contact_other *Contact, contact_own *Contact, data []byte) {
	ip, port := getIpPort(contact_own.Address)
	fmt.Println(ip, port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}

	conn, err := dialer.Dial("tcp", contact_other.Address)
	if err != nil {
		fmt.Println("Error caught: ", err)
		defer conn.Close()
	}

	fmt.Println("Connection was established to: ", conn.RemoteAddr())

	conn.Write([]byte("join"))
	tmp := make([]byte, 1024)
	time.Sleep(3 * time.Second)
	n, err := conn.Read(tmp)
	fmt.Println(n)
	receivedString := string(tmp[:n])
	fmt.Println(receivedString)
}

func Listen(ip string, port int, numberofreplicas *int, rt *RoutingTable) {
	ln, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))

	if err != nil {
		fmt.Println("Caught error: ", err)
		return
	}

	defer ln.Close()

	fmt.Println("Listening on ip and port", ip, port)

	for {
		// Accept incoming connections
		fmt.Println("HELLO NUMBER 1")
		conn, err := ln.Accept()
		fmt.Println("HELLO NUMBER 2")
		if err != nil {
			// handle error
			fmt.Println("Error:", err)
			continue
		}
		// Handle client connection in a goroutine
		//handleConnection(conn, numberofreplicas, rt)
		go RPC_handler(conn, rt)
	}

}

func handleConnection(conn net.Conn, numberofreplicas *int, rt *RoutingTable) {
	// handle incoming messages here
	fmt.Println("Connection accepted from", conn.RemoteAddr().String())

	*numberofreplicas += 1
	fmt.Println("Number of replicas: ", *numberofreplicas)

	address := conn.RemoteAddr().String()
	id := NewKademliaID(generateHashforNode(address))
	contact := NewContact(id, address)

	rt.AddContact(contact)

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println("In for loop contacts: ", contacts[i].String())
	}

	tmp := make([]byte, 1024)
	n, err := conn.Read(tmp)

	if err != nil {
		fmt.Println("Error caught: ", err)
		defer conn.Close()
	} else {
		// Convert the bytes read to a string
		receivedString := string(tmp[:n])
		fmt.Println("Received from connection:", receivedString)
		return_rt := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
		for i := range return_rt {
			fmt.Println(return_rt[i].Address)
			fmt.Println("hello")
			if return_rt[i].Address == "172.16.238.10:8080" || rt.me.Address == "172.16.238.10:8080" {

				fmt.Println("hello")
				conn.Write([]byte("joined"))
				break
			}
		}
	}

	//return conn
	defer conn.Close()
}

func SendPingMessage(contact_root *Contact, contact_own *Contact) {

	ip, port := getIpPort(contact_own.Address)
	fmt.Println(ip, port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}

	conn, err := dialer.Dial("tcp", contact_root.Address)
	if err != nil {
		fmt.Println("Error caught: ", err)
		defer conn.Close()
	}

	fmt.Println("Connection was established to: ", conn.RemoteAddr())
}

func getIpPort(address string) (ip string, port int) {
	port_number, err := strconv.Atoi(address[len(address)-4:])
	UNUSED(err)
	ip_address := address[:len(address)-5]
	fmt.Println("port number: ", port_number, "ip_address: ", ip_address)
	return ip_address, port_number
}

func Join(dst_address string, rt *RoutingTable) {
	address := returnIpAddress()

	ip, port := getIpPort(address)
	fmt.Println("address: ", address, " Port: ", port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}

	conn, err := dialer.Dial("tcp", dst_address)

	if err != nil {
		fmt.Println("Error caught: ", err)
		defer conn.Close()

	} else {
		fmt.Println("Connection established to: ", conn.RemoteAddr().String())

		conn.Write([]byte("join"))
		tmp := make([]byte, 1024)
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)
		fmt.Println(n)
		receivedString := string(tmp[:n])
		fmt.Println("Received from connection:", receivedString)
		if receivedString == "joined" {
			id_Root_Node := NewKademliaID(generateHashForRootNode())

			//generate a contact to the rootnode
			contact_RootNode := NewContact(id_Root_Node, "172.16.238.10:8080")
			UNUSED(contact_RootNode)
			rt.AddContact(contact_RootNode)
		}

		UNUSED(err)
		//
		return_rt := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
		for i := range return_rt {
			fmt.Println("CURRENT CONTACTS: ", return_rt[i].Address)
			fmt.Println("hello")
		}

		defer conn.Close()
	}

}

// this is our FindNode()
func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
