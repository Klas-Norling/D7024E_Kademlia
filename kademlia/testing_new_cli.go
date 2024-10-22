package kademlia

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CLIFORNODES() {
	// Define the file path
	filePath := "file.txt"

	// Step 1: Read the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Step 2: Split the file content into words and print each word
	words := strings.Fields(content) // Fields splits the content by spaces and newlines
	for _, word := range words {
		fmt.Println(word)
	}

	// Step 3: Overwrite the file with a space
	err = ioutil.WriteFile(filePath, []byte(" "), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File contents replaced with a space.")
}

func Clifornode() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("cli> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "hello":
			fmt.Println("Hello, User!")
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command:", input)
		}
	}
}
