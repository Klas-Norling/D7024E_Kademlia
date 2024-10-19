package kademlia

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Network struct {
}

type netnet struct {
	ditance *KademliaID
}

func MakeSenseOfStringMessage(recieved string) (string, string) {
	// Split the string by the semicolon
	parts := strings.SplitN(recieved, ";", 2)

	// If there are two parts, return them
	if len(parts) == 2 {
		return parts[0], parts[1]
	}

	// If the semicolon is not found, return the original string and an empty string
	return parts[0], ""

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

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func DecodeToPerson(s []byte) []Contact {

	p := []Contact{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	return p
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
	fmt.Println("Listen Listening on ip and port", ip, port)
	ln, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	//fmt.Println("HELOOOOOOOOOOO1")
	if err != nil {
		fmt.Println("Caught error: ", err)
		return
	}
	defer ln.Close()
	// Accept incoming connections
	//fmt.Println("HELOOOOOOOOOOO")
	for {
		//	fmt.Println("HELLO1")
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("Error:", err)
		}
		//	fmt.Println("HELLO2")
		//go RPC_handler(conn, rt)
		go RPC_handler(conn, rt)

		fmt.Println("Listen Listening on ip and port", ip, port)
		//	time.Sleep(3 * time.Second)
		UNUSED(conn)
	}

}

func RPC_handler(conn net.Conn, rt *RoutingTable) {
	//

	tmp := make([]byte, 1024)
	n, err := conn.Read(tmp)
	//root_id := NewKademliaID(generateHashForRootNode())
	//contact := NewContact(root_id, "172.16.238.10:8080")
	//var kadem = Kademlia{}
	//var network_struct = Network{}

	UNUSED(err)

	receivedString := string(tmp[:n])
	fmt.Println(receivedString)

	command, ipaddr := MakeSenseOfStringMessage(receivedString)
	fmt.Println(command + ipaddr)
	fmt.Println("command:", command, "other", ipaddr)
	switch command {
	case "store":
		// Initialize or reset the store
		kademlia := InitializeNode()

		address := conn.RemoteAddr().String()
		id := NewKademliaID(address)
		contact := NewContact(id, address)
		rt.AddContact(contact)
		kademlia.Store([]byte(ipaddr))

	case "find_node":
		// Get value by key

		newip_forsender, newport_forsender := getdecrementIpPort(conn.RemoteAddr().String())
		contacts, contact := switch_case_find_node(newip_forsender, ipaddr, newport_forsender, rt)
		/*
			address_forsender := newip_forsender + ":" + strconv.Itoa(newport_forsender)

			//we generate a kademliaID for the sender
			id_forsender := NewKademliaID(generateHashforNode(address_forsender))
			contact_forsender := NewContact(id_forsender, address_forsender)
			rt.AddContact(contact_forsender)

			newip, newport := getIpPort(ipaddr)
			//fmt.Println("Hello1234: ", newip+":"+strconv.Itoa(newport))
			address := newip + ":" + strconv.Itoa(newport)
			//fmt.Println("Hello1234: ", address)

			id := NewKademliaID(generateHashforNode(address))
			contact := NewContact(id, address)
			rt.AddContact(contact)

			//network_struct.SendFindContactMessage(&contact) //Useless???
			//fmt.Println("HHHHHHHHHHHHHHHHHHHHH", contact.ID)
			contacts := rt.FindClosestContacts(contact.ID, 20)*/
		//fmt.Println("HELELELLE", contacts[0].distance)
		//fmt.Println("HELELELLE", *contacts[0].distance)
		//fmt.Println(contacts)
		//bytesof_contacts, err := encodeContactsToBytes(contacts)
		//fmt.Println("HELLO!23")
		//NewSenderFunc(conn, &contact, &rt.me, bytesof_contacts)

		NewSenderFunc(conn, &contact, &rt.me, contacts)
		UNUSED(err)

	case "ping":
		fmt.Println("Connection was established to: ", conn.RemoteAddr())
		sendstring := []byte("pong")
		conn.Write(sendstring)

	case "find_value":

		address := conn.RemoteAddr().String()
		data := switch_case_find_value(address, ipaddr, rt)
		/*
			kademlia := InitializeNode()
			id := NewKademliaID(address)
			contact := NewContact(id, address)
			rt.AddContact(contact)
			data := EncodeToBytes(kademlia.LookupData(ipaddr))*/
		conn.Write(data)
		defer conn.Close()

	default:
		fmt.Println("RPC HANDLER DEFAULT")
		address := conn.RemoteAddr().String()
		id := NewKademliaID(generateHashforNode(address))
		contact := NewContact(id, address)
		rt.AddContact(contact)

	}

}

func switch_case_find_node(newip_forsender string, ipaddr string, newport_forsender int, rt *RoutingTable) ([]Contact, Contact) {
	address_forsender := newip_forsender + ":" + strconv.Itoa(newport_forsender)

	//we generate a kademliaID for the sender
	id_forsender := NewKademliaID(generateHashforNode(address_forsender))
	contact_forsender := NewContact(id_forsender, address_forsender)
	rt.AddContact(contact_forsender)

	newip, newport := getIpPort(ipaddr)
	//fmt.Println("Hello1234: ", newip+":"+strconv.Itoa(newport))
	address := newip + ":" + strconv.Itoa(newport)
	//fmt.Println("Hello1234: ", address)

	//
	id := NewKademliaID(generateHashforNode(address))
	contact := NewContact(id, address)
	rt.AddContact(contact)
	contacts := rt.FindClosestContacts(contact.ID, 20)
	fmt.Println("HELELELLELELELELELLELYEAH")
	return contacts, contact
}

func switch_case_find_value(address string, ipaddr string, rt *RoutingTable) []byte {
	kademlia := InitializeNode()
	//id := NewKademliaID(address)
	id := NewKademliaID(generateHashforNode(address))
	contact := NewContact(id, address)
	rt.AddContact(contact)
	data := EncodeToBytes(kademlia.LookupData(ipaddr))
	return data

}

func NewSenderFunc(conn net.Conn, contact_other *Contact, contact_own *Contact, contacts []Contact) {
	fmt.Println("Connection was established to: ", conn.RemoteAddr())
	// Create a gob encoder for the connection
	//encoder := gob.NewEncoder(conn)

	// Encode the struct array and send it over the TCP connection
	//err := encoder.Encode(contacts)
	//UNUSED(err)

	data := EncodeToBytes(contacts)
	//fmt.Println(contacts)
	//data2 := DecodeToPerson(data)
	//fmt.Println("IS DECODE WRONG???", data2)
	conn.Write(data)
}

func InitiateSender(dst_address string, data []byte, rt *RoutingTable, c chan []Contact) {
	address := returnIpAddress()

	ip, port := getNewIpPort(address)
	fmt.Println("in initiate sender address: ", address, " Port: ", port)
	return_contacts := []Contact{}
	fmt.Println("hello")
	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}
	fmt.Println("hello")
	conn, err := dialer.Dial("tcp", dst_address)
	fmt.Println("hello")
	if err != nil {
		fmt.Println("Error caught: ", err)
		defer conn.Close()

	} else {
		conn.Write(data)
		tmp := make([]byte, 2048)
		fmt.Println("hello1")
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)

		receivedString := string(tmp[:n])
		fmt.Println("hello2")

		recstring := []byte(receivedString)

		new_recstring := DecodeToPerson(recstring)
		fmt.Println("HLLO")
		UNUSED(err)

		return_contacts = new_recstring
		fmt.Println("HLLO2")
		c <- return_contacts
		fmt.Println("HLLO1")

	}
	defer conn.Close()
	fmt.Println("HELLO", return_contacts)
	fmt.Println("HELLO???")

}

func InitiateSenderForPong(dst_address string, data []byte, rt *RoutingTable, c chan string) {
	address := returnIpAddress()

	ip, port := getNewIpPort(address)
	fmt.Println("in initiate sender address: ", address, " Port: ", port)

	fmt.Println("hello")
	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}
	fmt.Println("hello")
	conn, err := dialer.Dial("tcp", dst_address)
	fmt.Println("hello")
	if err != nil {
		fmt.Println("Error caught: ", err)
		defer conn.Close()

	} else {
		conn.Write(data)
		tmp := make([]byte, 2048)
		fmt.Println("hello1")
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)

		receivedString := string(tmp[:n])
		fmt.Println("hello2")

		fmt.Println("HLLO")
		UNUSED(err)

		return_contacts := receivedString
		fmt.Println("HLLO2")
		c <- return_contacts
		fmt.Println("HLLO1")

	}
	defer conn.Close()

	fmt.Println("HELLO???")

}

/*
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
		RPC_handler(conn, rt)
	}

}*/
/*
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
}*/

func SendPingMessage(contact_root *Contact, contact_own *Contact) {

	ip, port := getNewNEWIpPort(contact_own.Address)
	fmt.Println(ip, port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}
	fmt.Println("root_ ", contact_root.Address)
	fmt.Println("own_ ", contact_own.Address)
	conn, err := dialer.Dial("tcp", contact_root.Address)
	if err != nil {
		fmt.Println("Error caught: ", err)

	}
	defer conn.Close()
	fmt.Println("Connection was established to---: ", conn.RemoteAddr())
}

func getIpPort(address string) (ip string, port int) {
	port_number, err := strconv.Atoi(address[len(address)-4:])
	UNUSED(err)
	ip_address := address[:len(address)-5]
	fmt.Println("port number: ", port_number, "ip_address: ", ip_address)
	return ip_address, port_number
}

func getdecrementIpPort(address string) (ip string, port int) {
	port_number, err := strconv.Atoi(address[len(address)-4:])
	new_portnumber := port_number - 1
	UNUSED(err)
	ip_address := address[:len(address)-5]
	fmt.Println("--port number: ", port_number, "--ip_address: ", ip_address)
	return ip_address, new_portnumber
}

func getNewIpPort(address string) (ip string, port int) {
	port_number, err := strconv.Atoi(address[len(address)-4:])
	new_port_number := port_number + 1
	UNUSED(err)
	ip_address := address[:len(address)-5]
	fmt.Println("port number: ", new_port_number, "ip_address: ", ip_address)
	return ip_address, new_port_number
}

func getNewNEWIpPort(address string) (ip string, port int) {
	port_number, err := strconv.Atoi(address[len(address)-4:])
	new_port_number := port_number + 2
	UNUSED(err)
	ip_address := address[:len(address)-5]
	fmt.Println("port number: ", new_port_number, "ip_address: ", ip_address)
	return ip_address, new_port_number
}

/*
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
*/
// this is our FindNode()
func (network *Network) SendFindContactMessage(contact Contact) {
	address := returnIpAddress()

	ip, port := getIpPort(address)
	fmt.Println("address: ", address, " Port: ", port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}
	conn, err := dialer.Dial("tcp", contact.Address)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		encode := EncodeToBytes("find_node;" + contact.Address)
		conn.Write(encode)
		tmp := make([]byte, 1024)
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)
		UNUSED(n, err)
		// receivedString := string(tmp[:n])

	}
}

func (network *Network) SendFindDataMessage(hash string, contact Contact) string {
	address := returnIpAddress()

	ip, port := getIpPort(address)
	fmt.Println("address: ", address, " Port: ", port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}

	conn, err := dialer.Dial("tcp", contact.Address)

	if err != nil {
		fmt.Println("Error: ", err)
		defer conn.Close()
	} else {
		encode := EncodeToBytes("find_value;" + hash)
		conn.Write(encode)
		tmp := make([]byte, 1024)
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)
		UNUSED(err)
		receivedString := string(tmp[:n])

		return receivedString

	}

	return "Error"

}

func (network *Network) SendStoreMessage(data string, contact Contact) {
	address := returnIpAddress()

	ip, port := getIpPort(address)
	fmt.Println("address: ", address, " Port: ", port)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}
	conn, err := dialer.Dial("tcp", contact.Address)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		encode := EncodeToBytes("store;" + data)
		conn.Write(encode)
		tmp := make([]byte, 1024)
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)
		UNUSED(n, err)
		// receivedString := string(tmp[:n])

	}
}
