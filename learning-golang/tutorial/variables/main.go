package main

import (
	"fmt"
)

// implicit declaration is not allowed outside of any method.
// yearsOfExperience := 7    //not allowed
// var yearsOfExperience int = 7  //allowed

// If it starts with a capital alphabet its a public variable.
const LoginToken string = "imcnjkoeicdh347977139478287mdvnfjelw"

func main() {
	var firstName string = "Kaustubh"
	var secondName string = "Desai"
	fmt.Println(firstName, secondName)
	fmt.Printf("Variable firstName is of type: %T \n", firstName)

	var smallVal uint8 = 255
	fmt.Println("smallVal", smallVal)

	var autoAdjust int = 535299525
	fmt.Println("autoAdjust", autoAdjust)
	fmt.Printf("Variable autoAdjust is of type: %T \n", autoAdjust)

	// by default value is 0
	var unassignedValue int
	fmt.Println("unassignedValue", unassignedValue)

	// sugar syntax/walrus operator
	yearsOfExperience := 7
	fmt.Println(yearsOfExperience)

	fmt.Println(LoginToken)
}
