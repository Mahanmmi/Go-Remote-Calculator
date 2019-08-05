package main

import (
	"net"
	// "../../calculator"
	"bufio"
	"fmt"
	"os"
)

func main() {
	go InitializeServer()
	fmt.Println("Main Hello")
	client, _ := net.Dial("tcp", "127.0.0.1:18757")
	fmt.Println("Main client", client)

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(client, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(client).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
