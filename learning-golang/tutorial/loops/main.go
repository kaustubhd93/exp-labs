package main

import (
	"fmt"
)

func main() {
	welcomeMessage := "############## Loops in go. ##############"
	fmt.Println(welcomeMessage)

	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// c style loop def
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, value := range days {
	// 	fmt.Println(value)
	// }

	initValue := 1
	for initValue < 10 {
		fmt.Println(initValue)
		if initValue == 4 {
			goto henlo
		}
		// execution flow jumped above, will not get back here.
		if initValue == 5 {
			break
		}
		initValue++
		// There is no such thing as ++initValue in go.
	}

henlo:
	fmt.Println("Goto block called")
}
