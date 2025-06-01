package main

import (
	"fmt"
)

func main() {
	welcomeMessage := "############## if else in go. ##############"
	fmt.Println(welcomeMessage)

	fmt.Println("Checking login count for user")
	loginCount := 10
	var message string
	if loginCount > 10 {
		message = "credentials probably shared, multiple devices login"
	} else if loginCount < 10 {
		message = "regular user"
	} else {
		message = "logins well within limit"
	}
	fmt.Println(message)

	if 11%2 == 0 {
		fmt.Println("Num is an even number")
	} else {
		fmt.Println("Num is an odd number")
	}

	if testNumber := 9; testNumber < 10 {
		fmt.Println("testNumber less than 10")
	}
}
