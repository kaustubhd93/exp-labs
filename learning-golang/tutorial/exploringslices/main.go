package main

import (
	"fmt"
	"sort"
)

func main() {
	welcomeMessage := "############## Slices in go. ##############"
	fmt.Println(welcomeMessage)

	var fruits = []string{"Apple", "Mango", "Blueberry"}
	fmt.Printf("Datatype of fruits: %T\n", fruits)

	fruits = append(fruits, "Banana", "Pineapple")
	fmt.Println("Fruits:", fruits)
	fmt.Println("3rd fruit:", fruits[2])

	// you can slice a slice. The first value is always inclusive and the last one is non inclusive.
	fruits = append(fruits[1:3])
	fmt.Println("Fruits:", fruits)
	fmt.Println("Fruits:", fruits)

	fruits = append(fruits[:3])
	fmt.Println("Fruits:", fruits)

	//make takes 2 arguments, the datatype and it's size
	testScores := make([]int, 4)
	testScores[0] = 34
	testScores[1] = 67
	testScores[2] = 78
	testScores[3] = 26
	// This below will give a runtime error. No syntax error will be thrown
	// it gives an error because we are trying to add a value where there is no memory allocation done.
	//testScores[4] = 89

	// However this append function is able to allocate memory when you want to expand the the slice
	testScores = append(testScores, 45, 77, 90)

	fmt.Println("Testscores:", testScores)
	fmt.Println("Are testScores sorted:", sort.IntsAreSorted(testScores))
	sort.Ints(testScores)
	fmt.Println("Testscores sorted:", testScores)
	fmt.Println("Are testScores sorted:", sort.IntsAreSorted(testScores))

	courses := []string{"ansible", "docker", "kubernetes", "jenkins", "terraform"}
	fmt.Println("courses:", courses)
	index := 2
	// removing an element in golang is not as straightforward as python where you can just pop the element
	// out of the list. Below syntax looks like breaking the list and joining them again.
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses:", courses)
}
