// Declare a package main
package main

import "fmt"

// Import the standard library fmt which is used for formatting text and
// printing info to the console.

// Implement a main function. A main function executes by default when you
// run code in the file.
// Also there can only be one main function in the entire project.
// Adding another main function in some other go file will result
// in an error saying it's already defined in Main.go
// func main() {
// 	//fmt.Println("Hello there...")
// 	var i int = 100
// 	i = 90
// 	// No need to give a data type when below type of declaration is used.
// 	j := "yo, what up?"
// 	var k float32 = 8.09
// 	// In below case 8.09 is inferred as float64
// 	// k := 8.09
// 	fmt.Println(i)
// 	fmt.Println(j)
// 	fmt.Printf("%v, %T\n", k, k)
// }

// func main() {
// 	// Everytime you initialise a variable and do not assign a value to it
// 	// it is initialised as 0 by default.
// 	// In case of a boolean it will take the value false.
// 	var flag bool
// 	fmt.Printf("%v,%T\n", flag, flag)
// 	m := 1 == 1
// 	n := 3 == 4
// 	fmt.Printf("%v,%T\n", m, m)
// 	fmt.Printf("%v,%T\n", n, n)
// }

// func main() {
// 	p := 43
// 	fmt.Printf("%v,%T\n", p, p)
// 	var q uint16 = 45678
// 	fmt.Printf("%v,%T\n", q, q)
// }

// func main() {
// 	// Check
// 	a := 10
// 	b := 6
// 	fmt.Printf("%v,%T\n", a, a)
// 	fmt.Printf("%v,%T\n", b, b)
// 	fmt.Println(a + b)
// 	fmt.Println(a * b)
// 	fmt.Println(a - b)
// 	fmt.Println(a / b)
// 	fmt.Println(a % b)
// 	var c int = 10
// 	var d int8 = 70
// 	// Cannot directly add int and int8 even when they are both integers.
// 	fmt.Println(c + int(d))
// }

// func main() {
// 	a := 11 // 1011
// 	b := 3  // 0011
// 	fmt.Println(a & b)
// 	fmt.Println(a | b)
// 	fmt.Println(a ^ b)
// 	// a 1 bit wherever the bit is set in a and not set in b.
// 	fmt.Println(a &^ b)
// }

// func main() {
// 	a := 56.6
// 	b := 11.3
// 	fmt.Println(a + b)
// 	fmt.Println(a * b)
// 	fmt.Println(a - b)
// 	fmt.Println(a / b)
// 	c := 4.9
// 	var d float32 = 4.9
// 	// Below operating is not allowed. => mismatched types float64 and float32
// 	// result := c == d
// 	result := float32(c) == d
// 	fmt.Println(result)
// }

// func main() {
// 	a := 1 + 2i
// 	b := 2 + 5.4i
// 	fmt.Println(a + b)
// 	fmt.Println(a * b)
// 	fmt.Println(a - b)
// 	fmt.Println(a / b)
//}

// func main() {
// 	s := "Hello world!"
// 	// s[n] will give the corresponding ascii valie of the character.
// 	fmt.Printf("%v, %T\n", s[1], s[1])
// 	// Need to typecast it for displaying the character.
// 	fmt.Printf("%v,%T\n", string(s[1]), string(s[1]))
// 	q := "I am Go!"
// 	// Easy concatenation like Python.
// 	fmt.Println(s + q)
// 	// Can be parsed as a byte slice to functions.
// 	b := []byte(s)
// 	fmt.Printf("%v,%T\n", b, b)
// 	// A string is declared in double quotes.
// 	// A rune is declated in single quotes.
// 	// rune is an alias for int32 and is equivalent
// 	// to int32 in all ways. It is used, by convention,
// 	// to distinguish character values from integer values.
// 	spc := '1'
// 	fmt.Printf("%v,%T\n", spc, spc)
// }

// Local const will be preferred if defined inside the function.
// const age int16 = 27

// func main() {
// 	// const a int = 45
// 	// var b int16 = 78
// 	const a = 45
// 	var b int16 = 78
// 	// Cannot evaluate a function or an expression and assign it as a constant.
// 	// const expn float64 = math.Sin(1.5)
// 	// Cannot reassign a const value
// 	// age = 67
// 	// Storing it in a variable and then assigning it to a constant
// 	// still wont work. It has to be a value.
// 	// var myExp float64 = math.Sin(1.5)
// 	// const finExp float64 = myExp

// 	// Below line does not give an error.
// 	// What the compiler does is, it uses a as a constant and does not consider
// 	// it's type because it sees 45 + b and not a+b.
// 	// It just adds and the final result is the data type of b
// 	fmt.Printf("%v, %T\n", a+b, a+b)
// 	fmt.Printf("%v, %T\n", a, a)
// }

// const (
// 	a = iota
// 	b
// 	c
// )

// const (
// 	d = iota
// 	e
// )

// func main() {
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)

// }

// func main() {
// grades := [4]int{56, 67, 90, 46}
// We can just add 3 dots in the place where we mention the size of the array at
// initiliasation and the compiler will automatically understand the size depending upon
// the no of elements we add in an array.
// grades := [...]int{56, 67, 90, 46}
// fmt.Printf("Grades : %v, %T\n", grades, grades)

// var students [4]string
// students[0] = "Kaustubh"
// students[1] = "Sagar"
// students[2] = "Anagha"
// students[3] = "James"
// fmt.Printf("Students : %v, %T\n", students, students)
// fmt.Printf("Student #4: %v\n", students[3])
// fmt.Printf("No of students : %v\n", len(students))

// a := [...]int{45, 67}
// b := a
// adding an & makes b point to a's memory address and directly
// make the change. If & is not used, b gets a copy of a and not
// the actual value.
// A typical example of value vs reference
// b := &a
// b[1] = 78
// fmt.Printf("a : %v\n", a)
// fmt.Printf("b : %v\n", b)

// a below here is a slice
// 	a := []int{45, 67}
// 	// var b []int
// 	b := a
// 	b[1] = 89
// 	fmt.Printf("a : %v, %T\n", a, a)
// 	fmt.Printf("Length of a : %v\n", len(a))
// 	fmt.Printf("Capacity of a : %v\n", cap(a))
// 	fmt.Printf("Capacity of b : %v\n", cap(b))
// 	fmt.Printf("a : %v, %T\n", a, a)
// }

// func main() {
// 	// a here is a slice.
// 	// a[<this element and elements after this>:<elements before this>]
// 	// a[3:6] => 4th,5th,6th
// 	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	b := a[:]   // slice of all elements
// 	c := a[3:]  // from 4th element to the end. 4th element inclusive
// 	d := a[:6]  // everything before the 7th element or first 6 elements
// 	e := a[3:6] // slice from 4th
// 	// "a" in below case will be altered because "b" is actually a slice of the array "a"
// 	// slices point directly to the original array and not their copies.
// 	b[5] = 42
// 	c[1] = 0
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)

// }

// Example of make function in go
func main() {
	//a := make([]int, 4)
	// a := make([]int, 4, 100)
	// fmt.Println(a)
	// fmt.Printf("length of a : %v\n", len(a))
	// fmt.Printf("capacity of a : %v\n", cap(a))
	a := []int{}
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
	a = append(a, 2)
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
	a = append(a, 3)
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
	a = append(a, 4)
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
	a = append(a, 5)
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
	a = append(a, 6, 7, 8, 9)
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
	// to add a slice of elements as individual elements
	// adding 3 dots at the end of the slice makes them individual elements
	a = append(a, []int{45, 67, 88, 100}...)
	fmt.Println(a)
	fmt.Printf("length of a : %v\n", len(a))
	fmt.Printf("capacity of a : %v\n", cap(a))
}
