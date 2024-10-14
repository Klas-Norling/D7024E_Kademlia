package kademlia

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Cli() {

	fmt.Println("Enter commands")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	inputHandler(text)

}

func inputHandler(input string) {
	fmt.Println("Input:" + input)

	switch input {
	case "get":
		fmt.Println("Getter")
	case "put":
		fmt.Println("Putter")

	default:
		fmt.Println("Default")
	}
}
