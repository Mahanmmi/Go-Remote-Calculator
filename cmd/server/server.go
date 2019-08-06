package main

import (
	"github.com/Mahanmmi/Go-Remote-Calculator/calculator"
	"encoding/gob"
	"github.com/bobesa/chalk"
	"fmt"
	"net"
	"strconv"
)

func clientHandler(clientSocket net.Conn) {
	dec := gob.NewDecoder(clientSocket)
	ao := &calculator.ArithmaticOperation{}
	dec.Decode(ao)
	chalk.Green().Bold().WhiteBackground().Printf("Received : %+v\n", ao)

	var ans float64
	switch ao.OperationType[:len(ao.OperationType)-2] {
	case "sum":
		{
			ans = calculator.Sum(ao.A, ao.B)
		}
	case "subtract":
		{
			ans = calculator.Subtract(ao.A, ao.B)
		}
	case "multiply":
		{
			ans = calculator.Multiply(ao.A, ao.B)
		}
	case "divide":
		{
			ans = calculator.Divide(ao.A, ao.B)
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
