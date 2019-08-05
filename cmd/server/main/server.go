package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func clientHandler(clientSocket net.Conn) {

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(clientSocket).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		clientSocket.Write([]byte(newmessage + "\n"))
	}
}

// InitializeServer TCP on port 18757
func InitializeServer() {
	fmt.Println("Server Hello")
	serverSocket, _ := net.Listen("tcp", "127.0.0.1:18757")
	for {
		client, _ := serverSocket.Accept()
		go clientHandler(client)
	}
}
