package kademlia

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Cli(rt *RoutingTable) {

	fmt.Println("Enter commands")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	inputHandler(text, rt)
}

func inputHandler(input string, rt *RoutingTable) {
	fmt.Println("Input:" + input)

	split := strings.SplitN(input, " ", 2)
	command := split[0]
	content := split[1]

	var network = Network{}
	var kademlia = Kademlia{}

	switch command {

	case "put":
		fmt.Println("Putter")
		fmt.Println("Content: ", content)

		kademlia.Store([]byte(content))

		// hash := kad.Store([]byte(content))
		// fmt.Println(hash)

		contacts := rt.FindClosestContacts(rt.me.ID, 3)

		for i := 0; i < len(contacts); i++ {
			network.SendStoreMessage([]byte(content), contacts[i])
		}

	case "get":
		fmt.Println("Putter")
		fmt.Println("Content: ", content)

		kademlia.LookupData(content)

		// value := kad.LookupData(content)
		// fmt.Println(value)

	case "exit":
		fmt.Println("Exiting")
		os.Exit(0)

	default:
		fmt.Println("Command not found")
	}
}
