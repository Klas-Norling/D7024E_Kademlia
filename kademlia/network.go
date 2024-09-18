package kademlia

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

type Network struct {
}

func Listen(ip string, port int) {
	ln, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))

	if err != nil {
		fmt.Println("Caught error: ", err)
		return
	}

	defer ln.Close()

	fmt.Println("Listening on ip and port", ip, port)

	for {
		// Accept incoming connections
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("Error:", err)
			continue
		}
		// Handle client connection in a goroutine
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn, array *[]string) (connection net.Conn) {
	// handle incoming messages here
	fmt.Println("Connection accepted from", conn.RemoteAddr().String())

	tmp := make([]byte, 1024)
	data := make([]byte, 0)
	length := 0
	n, err := conn.Read(tmp)

	// loop through the connection stream, appending tmp to data
	for {
		// read to the tmp var
		n, err := conn.Read(tmp)
		if err != nil {
			// log if not normal error
			if err != io.EOF {
				fmt.Printf("Read error - %s\n", err)
			}
			break
		}

		// append read data to full data
		data = append(data, tmp[:n]...)

		// update total read var
		length += n
	}

	fmt.Println(data)
	fmt.Println(n, err)

	// log bytes read
	fmt.Printf("READ  %d bytes\n", length)
	return conn
	// defer conn.Close()
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

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}

func UNUSED(x ...interface{}) {}
