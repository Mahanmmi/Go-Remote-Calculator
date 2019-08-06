package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
)

// ArithmaticOperation struct
type ArithmaticOperation struct {
	OperationType string
	A, B          float64
}

// InitializeNewClient :|
func InitializeNewClient() {
	client, _ := net.Dial("tcp", "127.0.0.1:18757")
	
	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Operation: (sum, subtract, multiply, divide): ")
	operation, _ := reader.ReadString('\n')
	fmt.Print("Enter first number: ")
	strNumber1, _ := reader.ReadString('\n')
	fmt.Print("Enter second number: ")
	strNumber2, _ := reader.ReadString('\n')

	// fmt.Println("s " + strNumber1 + " f")
	// fmt.Println("s " + strNumber2 + " f")

	number1, _ := strconv.ParseFloat(strNumber1[:len(strNumber1)-2], 64)
	number2, _ := strconv.ParseFloat(strNumber2[:len(strNumber2)-2], 64)

	// fmt.Println("s" , number1 , "f")
	// fmt.Println("s" , number2 , "f")
		
		
	// encode and send data
	data := ArithmaticOperation{operation, number1, number2}
	encoder := gob.NewEncoder(client)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}

	// listen for reply
	message, _ := bufio.NewReader(client).ReadString('\n')
	fmt.Print("Answer from server: " + message)
}
