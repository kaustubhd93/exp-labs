package main

import (
	"fmt"
)

func main() {
	// This a slice because it's size is not fixed.
	var taskItems = []string{"Learn basics of golang", "Develop a handy url checker cli tool", "Reward myself with cheesecake"}
	fmt.Println("##### Hey there! This is just another todoapp. #####")
	fmt.Println("Todo list for this week:")
	fmt.Println(taskItems)
}
