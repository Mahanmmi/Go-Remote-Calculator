package main

import (
	// "bufio"
	"../../calculator"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// ArithmaticOperation struct
type ArithmaticOperation struct {
	OperationType string
	A, B          float64
}

func clientHandler(clientSocket net.Conn) {
	dec := gob.NewDecoder(clientSocket)
	p := &ArithmaticOperation{}
	dec.Decode(p)
	fmt.Printf("Received : %+v", p)

	var ans float64
	switch p.OperationType[:len(p.OperationType)-2] {
	case "sum":
		{
			ans = calculator.Sum(p.A, p.B)
		}
	case "subtract":
		{
			ans = calculator.Subtract(p.A, p.B)
		}
	case "multiply":
		{
			ans = calculator.Multiply(p.A, p.B)
		}
	case "divide":
		{
			ans = calculator.Divide(p.A, p.B)
		}
	}
	strAns := strconv.FormatFloat(ans, 'f', 2, 64)
	clientSocket.Write([]byte(strAns + "\n"))
	clientSocket.Close()
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
