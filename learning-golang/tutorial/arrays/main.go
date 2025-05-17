package main

import (
	"fmt"
)

func main() {
	welcomeMessage := "############## Arrays in go ##############"
	fmt.Println(welcomeMessage)

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[2] = "Mango"
	fruits[3] = "Blueberry"

	fmt.Println("Fruits:", fruits)
	fmt.Println("Length of fruits array:", len(fruits))
	fmt.Printf("Datatype of fruits: %T\n", fruits)

	var veggies = [5]string{"Asparagus", "Okra", "Fenugreek"}
	fmt.Println(veggies)
	fmt.Printf("Datatype of veggies: %T\n", veggies)
	// veggies[1] = "Spinach"
	// fmt.Println(veggies)
}
