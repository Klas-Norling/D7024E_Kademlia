package kademlia

import (
	"fmt"
	"os"
	"time"
)

func Cli(kad *Kademlia) {
	//kad := InitializeNode()
	//fmt.Println("CLI RT", rt.FindClosestContacts(rt.me.ID, 3))
	//kad.rt = *rt
	for true {
		//fmt.Println("Enter commands")
		time.Sleep(time.Second * 1)
		words := CLIFORNODES()
		rt := kad.GetRoutingtable()
		inputHandler(words, &rt, kad)
	}
}

func inputHandler(input []string, rt *RoutingTable, kad *Kademlia) {

	//fmt.Println("hello gain CLI")
	//fmt.Println(input[0])
	//fmt.Println("hello gain CLI")

	switch input[0] {

	case "put":
		fmt.Println("Putter")
		fmt.Println("Content: ", input[1])
		store_in_bytes := []byte(input[1])
		fmt.Println(store_in_bytes)
		kad.Store(store_in_bytes)

		// hash := kad.Store([]byte(content))
		// fmt.Println(hash)

		contacts := rt.FindClosestContacts(rt.me.ID, 3)

		for i := 0; i < len(contacts); i++ {
			kad.network.SendStoreMessage(input[1], contacts[i])
		}
		UNUSED(contacts)

	case "get":
		fmt.Println("Getter")
		fmt.Println("Content: ", input[1])

		kad.LookupData(input[1])

		// value := kad.LookupData(content)
		// fmt.Println(value)

	case "exit":
		fmt.Println("Exiting")
		os.Exit(0)

	case "hello":

	default:
		fmt.Println("Command not found")
	}
}
