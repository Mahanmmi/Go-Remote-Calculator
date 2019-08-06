package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"github.com/Mahanmmi/Go-Remote-Calculator/calculator"
	"github.com/bobesa/chalk"
	"net"
	"os"
	"strconv"
)

// NewClient :|
func NewClient() {
	client, _ := net.Dial("tcp", "127.0.0.1:18757")

	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Operation: (sum, subtract, multiply, divide): ")
	operation, _ := reader.ReadString('\n')
	fmt.Print("Enter first number: ")
	strNumber1, _ := reader.ReadString('\n')
	fmt.Print("Enter second number: ")
	strNumber2, _ := reader.ReadString('\n')

	// Converting input string to numbers
	number1, _ := strconv.ParseFloat(strNumber1[:len(strNumber1)-2], 64)
	number2, _ := strconv.ParseFloat(strNumber2[:len(strNumber2)-2], 64)

	// encode and send data
	data := calculator.ArithmaticOperation{operation, number1, number2}
	encoder := gob.NewEncoder(client)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}

	// listen for reply
	message, _ := bufio.NewReader(client).ReadString('\n')

	chalk.GreenBackground().Bold().Black().Print("Answer from server: " + message)
}
