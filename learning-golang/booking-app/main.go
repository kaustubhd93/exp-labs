package main

import "fmt"

func main() {
	var conferenceName = "GOPHERCON 2024"
	const totalTickets = 100
	var ticketsAvailable = 100

	fmt.Println("**************************************************************")
	fmt.Println("Hey there! Welcome to", conferenceName, "booking application")
	fmt.Println("**************************************************************")
	fmt.Println("************ Get your tickets here ************")
	// %v will substitute the value of the variable.
	// %T will subsitute the data type of the variable.
	fmt.Printf("Total no of tickets: %v\n", totalTickets)
	fmt.Printf("Tickets available: %v\n", ticketsAvailable)

	var userName string
	var userTickets int
	userName = "john.doe"
	userTickets = 3
	fmt.Printf("Hi %v! you have booked %v tickets\n", userName, userTickets)

}
