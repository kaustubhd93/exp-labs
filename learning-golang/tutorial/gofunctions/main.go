package main

import (
	"fmt"
)

func main() {
	welcomeMessage := "############## Functions in go. ##############"
	fmt.Println(welcomeMessage)
	dummy()

	// not allowed
	// func dummy2() {
	// 	fmt.Println("Hello world inside")
	// }

	dummy2()
	fmt.Println(funcadd(4, 7))
	fmt.Println(dynamicAdd(56, 7, 9, 2, 3, 83))
}

func dummy() {
	fmt.Println("World says hello!")
}

func dummy2() {
	fmt.Println("Hello world inside")
}

func funcadd(x int, y int) int {
	return x + y
}

// When we dont know the number of arguments which are going to be passed.
// values here is inferred as a slice.
func dynamicAdd(values ...int) int {
	finalTotal := 0

	for _, val := range values {
		finalTotal += val
	}
	return finalTotal
}
