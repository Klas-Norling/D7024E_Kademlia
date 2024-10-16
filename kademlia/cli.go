package kademlia

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Cli(kademlia *Kademlia) {

	fmt.Println("Enter commands")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	//time.Sleep(time.Second * 10)
	InputHandler(text, kademlia)
}

func InputHandler(input string, kademlia *Kademlia) {
	fmt.Println("Input:" + input)

	split := strings.SplitN(input, " ", 2)
	command := split[0]
	content := split[1]

	var network = Network{}
	switch command {

	case "put":
		fmt.Println("Putter")
		fmt.Println("Content: ", content)

		kademlia.Store([]byte(content))

		fmt.Println("After store command")
		// hash := kad.Store([]byte(content))
		// fmt.Println(hash)

		contacts := kademlia.rt.FindClosestContacts(kademlia.rt.me.ID, 3)

		for i := 0; i < len(contacts); i++ {
			network.SendStoreMessage(content, contacts[i])
		}

	case "get":
		fmt.Println("Getter")
		fmt.Println("Content: ", content)

		//kademlia.LookupData(content)
		kademlia.LookupData("a1b9bdfcb1f469376df7431bbb2a375fb3fb413a")

		// value := kad.LookupData(content)
		// fmt.Println(value)

	case "exit":
		fmt.Println("Exiting")
		os.Exit(0)

	default:
		fmt.Println("Command not found")
	}
}
