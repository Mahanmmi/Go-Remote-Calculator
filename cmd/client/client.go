package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bobesa/chalk"
)

// NewClient :|
func NewClient() {
	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Operation: (sum, subtract, multiply, divide): ")
	operation, _ := reader.ReadString('\n')
	fmt.Print("Enter first number: ")
	strNumber1, _ := reader.ReadString('\n')
	fmt.Print("Enter second number: ")
	strNumber2, _ := reader.ReadString('\n')

	url := "http://127.0.0.1:18757/calculator" + "?type=" + operation[:len(operation)-2] + "&A=" + strNumber1[:len(strNumber1)-2] + "&B=" + strNumber2[:len(strNumber2)-2]

	fmt.Println(url)

	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	// listen for reply

	chalk.GreenBackground().Bold().Black().Print("Answer from server: " + string(bytes))
}
