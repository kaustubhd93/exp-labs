package main

import (
	"fmt"
	"net/http"
)

// This a slice because it's size is not fixed.
var taskItems = []string{"Learn basics of golang", "Develop a utility: url checker cli tool", "Reward myself with blueberry cheesecake"}

func main() {
	//fmt.Println("##### Hey there! This is just another todoapp. #####")
	// fmt.Println("Todo list for this week:")
	// printTasks(taskItems)
	// fmt.Println("******************************************")
	// taskItems = addTask(taskItems, "Go for a walk")
	// fmt.Println("Updated To do list.")
	// printTasks(taskItems)
	// fmt.Println("******************************************")

	// use _ for using only values in array.
	// for _, task := range taskItems {
	// 	fmt.Println(task)
	// }
	http.HandleFunc("/", homePage)
	http.HandleFunc("/tasks", showTasks)
	fmt.Println("Will start listening on port 8085")
	http.ListenAndServe(":8085", nil)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	welcomeMessage := "##### Hey there! This is just another todoapp. #####"
	fmt.Fprintln(writer, welcomeMessage)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Tasks for the week:")
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func printTasks(taskItems []string) {
	fmt.Println("Tasks: ")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTasks = append(taskItems, newTask)
	return updatedTasks
}
