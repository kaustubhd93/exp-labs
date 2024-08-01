package main

import (
	"fmt"
	"strings"
)

func main() {
	// Golang can infer the data type of the variable.
	var conferenceName = "GOPHERCON 2024"
	// Sometimes if golang doesn't infer you can statically type it.
	const totalTickets int = 100
	// Go's sugar syntax
	// ticketsAvailable := 100
	var ticketsAvailable uint = 100
	// While declaring an array size needs to be mentioned. This size is fixed.
	// var bookings [100]string
	// In case if we dont the size of array and it's index for storing data, we can use
	// slices. Slices are basically flexible arrays. It's effecient
	var bookings []string

	fmt.Println("**************************************************************")
	fmt.Println("Hey there! Welcome to", conferenceName, "booking application")
	fmt.Println("**************************************************************")
	fmt.Println("************ Get your tickets here ************")
	// %v will substitute the value of the variable.
	// %T will subsitute the data type of the variable.
	fmt.Printf("Total no of tickets: %v\n", totalTickets)
	fmt.Printf("Tickets available: %v\n", ticketsAvailable)

	for {
		var userName string
		var firstName string
		var lastName string
		var emailId string
		var userTickets uint
		fmt.Print("Enter your username: ")
		fmt.Scan(&userName)
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email id: ")
		fmt.Scan(&emailId)
		fmt.Print("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		// validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(emailId, "@") && strings.Contains(emailId, ".")
		isValidUserTickets := userTickets > 0 && userTickets <= ticketsAvailable

		// Check if all validations are met
		if !isValidName || !isValidEmail || !isValidUserTickets {
			if userTickets > ticketsAvailable {
				fmt.Println("*************************************************************")
				fmt.Printf("Only %v ticket(s) are available\n", ticketsAvailable)
				fmt.Println("*************************************************************")
			}
			fmt.Println("*************************************************************")
			fmt.Println("Unable to book ticket. Please enter valid details.")
			fmt.Println("*************************************************************")
			if !isValidName {
				fmt.Println("First Name and Last name should have atleast 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Email id entered is not a valid email id")
			}
			// Dont break the loop, give an option to book the right amount of tickets to user.
			continue
		}
		ticketsAvailable = ticketsAvailable - userTickets
		//bookings[0] = firstName + " " + lastName + " " + emailId
		bookings = append(bookings, firstName+" "+lastName+" "+emailId)

		fmt.Printf("Hi %v! you have booked %v tickets\n", firstName, userTickets)
		fmt.Printf("Tickets available: %v\n", ticketsAvailable)
		fmt.Printf("Total bookings: %v\n", bookings)
		firstNames := []string{}
		// for index, element_of_list := range slice
		// _ is used when we want to declare an unused variable. This is called a blank identifier.
		for _, bookingDetails := range bookings {
			// using the strings.Fields function will split the element by spaces and create a new list
			// that can be iterated. Works similar to "dummy_string".split() in python
			var nameDetails = strings.Fields(bookingDetails)
			firstNames = append(firstNames, nameDetails[0])
		}
		fmt.Printf("First name of people who have booked tickets: %v\n", firstNames)

		if ticketsAvailable == 0 {
			fmt.Println("All tickets sold out!")
			break
		}
	}

}
