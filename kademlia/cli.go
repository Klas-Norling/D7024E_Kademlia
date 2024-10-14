package kademlia

import (
	"bufio"
	"fmt"
	"os"
)

func Cli() {
	fmt.Println("Enter commands")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}
