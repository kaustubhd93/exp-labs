package main

import "fmt"

func main() {
	// Golang can infer the data type of the variable.
	var conferenceName = "GOPHERCON 2024"
	// Sometimes if golang doesnt infer you can statically type it.
	const totalTickets int = 100
	// Go's sugar syntax
	// ticketsAvailable := 100
	var ticketsAvailable uint = 100

	fmt.Println("**************************************************************")
	fmt.Println("Hey there! Welcome to", conferenceName, "booking application")
	fmt.Println("**************************************************************")
	fmt.Println("************ Get your tickets here ************")
	// %v will substitute the value of the variable.
	// %T will subsitute the data type of the variable.
	fmt.Printf("Total no of tickets: %v\n", totalTickets)
	fmt.Printf("Tickets available: %v\n", ticketsAvailable)

	var userName string
	var userTickets uint
	fmt.Print("Enter your username: ")
	fmt.Scan(&userName)
	fmt.Print("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	ticketsAvailable = ticketsAvailable - userTickets

	fmt.Printf("Hi %v! you have booked %v tickets\n", userName, userTickets)
	fmt.Printf("Tickets available: %v\n", ticketsAvailable)

}
