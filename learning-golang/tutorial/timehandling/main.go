package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.Now())
	presentTime := time.Now()
	fmt.Println("currentTime:", presentTime)
	// standard time format if you want to print datetime in a certain way.
	// cannot use values other than this.
	// source = https://pkg.go.dev/time#example-Time.Format
	fmt.Println("formatted time:", presentTime.Format("01-02-2006"))
	fmt.Println("formatted time:", presentTime.Format("01-02-2006 Monday"))
	fmt.Println("formatted time:", presentTime.Format("02-01-2006 15:04:05 Monday"))

	// creating any date.
	createdDate := time.Date(2019, time.March, 6, 4, 5, 2, 9, time.UTC)
	fmt.Println("Created date:", createdDate)
	fmt.Println("formatted time:", createdDate.Format("01-02-2006 Monday"))
}
