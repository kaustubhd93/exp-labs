package main

import (
	"fmt"
)

func main() {
	welcomeMessage := "Pointers in go."
	fmt.Println(welcomeMessage)

	var pntr *int
	fmt.Println("pntr:", pntr)

	sampleVal := 56
	actualPntr := &sampleVal
	sampleCopyVal := sampleVal
	// This will give the memory address where the value is stored
	fmt.Println("actualPntr:", actualPntr)
	// This till give the actual value of the variable.
	fmt.Println("actualPntr:", *actualPntr)

	*actualPntr = *actualPntr * 2
	// The original value which has been reference everywhere now has been modified.
	// pointers ensure the actual value is being referenced and operated no ops on copies.
	fmt.Println("Modified value:", sampleVal)
	sampleCopyVal = sampleCopyVal * 2
	fmt.Println("Modified value of copy:", sampleCopyVal)
	fmt.Println("Memory address of value of copy:", &sampleCopyVal)
}
