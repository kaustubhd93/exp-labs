package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcomeMessage := "Welcome to our Pizza app."
	fmt.Println(welcomeMessage)
	fmt.Println("What do you think of our Pizza on a scale of 1 to 5?")
	// use bufio to accept input from stdin
	reader := bufio.NewReader(os.Stdin)

	// comma ok || error err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating: ", input)
	// fmt.Println(err)

	// manipulate user rating
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Manipulated rating is : ", numRating+1)
	}

}
