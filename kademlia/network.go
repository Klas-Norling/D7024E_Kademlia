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
	fmt.Println("recieved: ", recieved)
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
		RPC_handler(conn, rt)

		fmt.Println("Listen Listening on ip and port", ip, port)
		//	time.Sleep(3 * time.Second)
		UNUSED(conn)
	}

}

func RPC_handler(conn net.Conn, rt *RoutingTable) {
	//
	old_address := conn.RemoteAddr().String()
	address, ip_port := getdecrementIpPort(old_address)
	address = address + ":" + strconv.Itoa(ip_port)
	id := NewKademliaID(generateHashforNode(address))
	contact := NewContact(id, address)
	rt.AddContact(contact)

	tmp := make([]byte, 1024)
	n, err := conn.Read(tmp)
	contactsrt := rt.FindClosestContacts(rt.me.ID, 3)
	fmt.Println("contactsrt: ", contactsrt)

	UNUSED(err)

	receivedString := string(tmp[:n])
	fmt.Println(receivedString)
	defer conn.Close()
	command, ipaddr := MakeSenseOfStringMessage(receivedString)
	fmt.Println(command + ipaddr)
	fmt.Println("command:", command, "other", ipaddr)
	switch command {
	case "store":
		// Initialize or reset the store
		fmt.Println("HELLO")
		kademlia := InitializeNode()

		address := conn.RemoteAddr().String()
		fmt.Println("ipaddr:" + ipaddr)

		id := NewKademliaID(generateHashforNode(address))
		fmt.Println("ipaddr:" + ipaddr)
		contact := NewContact(id, address)
		fmt.Println("ipaddr:" + ipaddr)
		rt.AddContact(contact)
		fmt.Println("ipaddr:" + ipaddr)
		kademlia.Store([]byte(ipaddr))
		conn.Close()

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
		conn.Close()

	case "ping":
		fmt.Println("Connection was established to: ", conn.RemoteAddr())
		sendstring := []byte("pong")
		conn.Write(sendstring)
		conn.Close()

	case "find_value":

		address := conn.RemoteAddr().String()
		data := switch_case_find_value(address, ipaddr, rt)
		fmt.Println("datavalue:", string(data))
		/*
			kademlia := InitializeNode()
			id := NewKademliaID(address)
			contact := NewContact(id, address)
			rt.AddContact(contact)
			data := EncodeToBytes(kademlia.LookupData(ipaddr))*/
		conn.Write(data)
		conn.Close()

	default:
		fmt.Println("RPC HANDLER DEFAULT")
		address := conn.RemoteAddr().String()
		id := NewKademliaID(generateHashforNode(address))
		contact := NewContact(id, address)
		rt.AddContact(contact)
		conn.Close()

	}

}

func switch_case_find_node(newip_forsender string, ipaddr string, newport_forsender int, rt *RoutingTable) ([]Contact, Contact) {
	/*
		address_forsender := newip_forsender + ":" + strconv.Itoa(newport_forsender)

		//we generate a kademliaID for the sender

		id_forsender := NewKademliaID(generateHashforNode(address_forsender))
		contact_forsender := NewContact(id_forsender, address_forsender)
		rt.AddContact(contact_forsender)*/

	newip, newport := getIpPort(ipaddr)
	//fmt.Println("Hello1234: ", newip+":"+strconv.Itoa(newport))
	address := newip + ":" + strconv.Itoa(newport)
	//fmt.Println("Hello1234: ", address)

	//
	id := NewKademliaID(generateHashforNode(address))
	contact := NewContact(id, address)
	//rt.AddContact(contact)
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
	data := kademlia.LookupData(generateHashforNode(ipaddr))
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
		fmt.Println("Error caught---: ", err)
		defer conn.Close()

	} else {
		fmt.Println(dst_address)
		conn.Write(data)
		tmp := make([]byte, 2048)
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)

		receivedString := string(tmp[:n])

		recstring := []byte(receivedString)

		new_recstring := DecodeToPerson(recstring)

		UNUSED(err)

		return_contacts = new_recstring

		c <- return_contacts

	}
	defer conn.Close()

}

/*
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
		fmt.Println(dst_address)
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

}*/

func SendPingMessage(contact_root *Contact, contact_own *Contact) {

	ip, port := getNewIpPort(contact_own.Address)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}

	conn, err := dialer.Dial("tcp", contact_root.Address)
	if err != nil {
		fmt.Println("Error caught: ", err)

	}
	encode := []byte("ping")
	conn.Write(encode)
	tmp := make([]byte, 1024)
	n, err := conn.Read(tmp)
	UNUSED(n, err)

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
	fmt.Println("pppppport number: ", port_number, "--ip_address: ", ip_address)

	return ip_address, new_portnumber
}

func GetdecrementIpPort(address string) (ip string, port int) {
	port_number, err := strconv.Atoi(address[len(address)-4:])
	fmt.Println("what does this print:", address[:len(address)-5])
	new_portnumber := port_number - 1
	UNUSED(err)
	ip_address := address[:len(address)-5]
	fmt.Println("pppppport number: ", port_number, "--ip_address: ", ip_address)

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

// this is our FindNode()
func (network *Network) SendFindContactMessage(contact Contact) {
	address := returnIpAddress()

	ip, port := getNewIpPort(address)

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

	ip, port := getNewIpPort(address)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}

	conn, err := dialer.Dial("tcp", contact.Address)

	if err != nil {
		fmt.Println("Error: ", err)
		conn.Close()
	} else {
		encode := []byte("find_value;" + hash)
		conn.Write(encode)
		tmp := make([]byte, 1024)
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)
		UNUSED(err)
		receivedString := string(tmp[:n])
		conn.Close()

		return receivedString

	}
	defer conn.Close()
	return "Error"

}

func (network *Network) SendStoreMessage(data string, contact Contact) {
	address := returnIpAddress()

	ip, port := getNewIpPort(address)

	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
	}
	conn, err := dialer.Dial("tcp", contact.Address)

	if err != nil {
		fmt.Println("Error: ", err)
		conn.Close()
	} else {
		encode := []byte("store;" + data)
		conn.Write(encode)
		tmp := make([]byte, 1024)
		time.Sleep(3 * time.Second)
		n, err := conn.Read(tmp)
		UNUSED(n, err)
		conn.Close()
		// receivedString := string(tmp[:n])

	}
	defer conn.Close()
}

//docker exec -it new_kadem-root_node-1 /bin/sh

//./command.sh file.txt "put 1234klas"
