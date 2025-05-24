package main

import (
	"fmt"
)

func main() {
	welcomeMessage := "############## Structs in go. ##############"
	fmt.Println(welcomeMessage)

	// there is no inheritance in go
	// no super or parent either.
	employeeDetails := Employee{"Barry Allen", 23, "Employed", "Male", "barry.allen@starlabs.com"}
	fmt.Println(employeeDetails)
	fmt.Printf("employeeDetails type: %T\n", employeeDetails)
	fmt.Printf("Complete details with key: %+v\n", employeeDetails)
	fmt.Printf("name of the employee: %v\nAge: %v\n", employeeDetails.Name, employeeDetails.Age)

}

type Employee struct {
	Name   string
	Age    int
	Status string
	Gender string
	Email  string
}
