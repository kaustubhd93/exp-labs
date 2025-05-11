package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "hello there! whats your task for today?"
	fmt.Println(welcomeMessage)
	// use bufio to accept input from stdin
	reader := bufio.NewReader(os.Stdin)

	// comma ok || error err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Tasks for today:", input)
	fmt.Printf("Type of input is : %T", input)
	// fmt.Println(err)
}
