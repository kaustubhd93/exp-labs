package main

import (
	"fmt"
)

type Employee struct {
	Name                string
	Age                 int
	Department          string
	EmploymentStatus    bool
	EmploymentStartDate string
	Email               string
}

// Methods are functions with a recievable argument
// primary use case is interfaces.
func (e Employee) getEmploymentStatus() bool {
	return e.EmploymentStatus
}

func (e Employee) updateEmploymentStatus() bool {
	// a copy of struct is being passed here
	e.EmploymentStatus = false
	return e.EmploymentStatus
}

func main() {
	welcomeMessage := "############## Methods in go. ##############"
	fmt.Println(welcomeMessage)
	ba := Employee{"Barry Allen", 24, "Special Forces", true, "29-09-2015", "barry.allen@starlabs.com"}
	fmt.Println(ba)
	fmt.Println(ba.getEmploymentStatus())
	// fmt.Println(ba.EmploymentStatus)
	fmt.Println(ba.updateEmploymentStatus())
	// original values are not affected by previous call.
	fmt.Println(ba)
	ba.EmploymentStatus = false
	// original value affected.
	fmt.Println(ba)
}
