// Tutorial Source : https://www.youtube.com/watch?v=YS4e4q9oBaU&t=9347s
// Declare a package main
package main

import (
	"fmt"
	"runtime"
	"sync"
	// "net/http"
	//"log"
	// "io/ioutil"
	// "reflect"
)

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
// func main() {
// 	//a := make([]int, 4)
// 	// a := make([]int, 4, 100)
// 	// fmt.Println(a)
// 	// fmt.Printf("length of a : %v\n", len(a))
// 	// fmt.Printf("capacity of a : %v\n", cap(a))
// 	a := []int{}
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// 	a = append(a, 1)
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// 	a = append(a, 2)
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// 	a = append(a, 3)
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// 	a = append(a, 4)
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// 	a = append(a, 5)
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// 	a = append(a, 6, 7, 8, 9)
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// 	// to add a slice of elements as individual elements
// 	// adding 3 dots at the end of the slice makes them individual elements
// 	a = append(a, []int{45, 67, 88, 100}...)
// 	fmt.Println(a)
// 	fmt.Printf("length of a : %v\n", len(a))
// 	fmt.Printf("capacity of a : %v\n", cap(a))
// }

// Timestamp: 2:37

// // Animal ...
// type Animal struct {
// 	Name   string
// 	Origin string
// }

// // Bird ...
// type Bird struct {
// 	// Bird has animal like characteristics like a name, an origin
// 	// By just writing the name of the struct here, we can borrow it's pieces.
// 	// This is not traditional inheritance but can be called as composition
// 	// To be short it can be called as a "has a" relationship.
// 	Animal
// 	SpeedKPH float32
// 	canFly   bool
// }

// func main() {
// 	fmt.Println("------Embedding in Structs-------")
// 	b := Bird{}
// 	b.Name = "Ostrich"
// 	b.Origin = "Africa"
// 	b.SpeedKPH = 70
// 	b.canFly = false
// 	fmt.Println(b)
// 	fmt.Println(b.Name)
// 	fmt.Println(b.canFly)
// 	// While using the literal struct syntax,
// 	// we will have to use the embedded struct in literal syntax.
// 	// If declared empty, we can use the embedded pieces directly.
// 	c := Bird{
// 		Animal:   Animal{Name: "Peacock", Origin: "India"},
// 		SpeedKPH: 5,
// 		canFly:   false,
// 	}
// 	fmt.Println(c)
// 	fmt.Println(c.Name)
// }

// // Animal ...
// type Animal struct {
// 	// If we wanted to validate the length of the characters in the name field.
// 	// we can use tags. It is the meta data we can use for validation.
// 	// Refer this https://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go
// 	Name   string `required_max: "100"`
// 	Origin string
// }

// func main() {
// 	fmt.Println("--------Tags in structs--------")
// 	stType := reflect.TypeOf(Animal{})
// 	field, _ := stType.FieldByName("Name")
// 	fmt.Println(field.Tag)
// 	// But what do we do with this tag? This will be decided by the vaildation
// 	// framework or logic we use. This tag value is meaningless to GO itself.
// }

// Timestamp : 2:48:03
// Control flow : if and switch statements
// func main() {
// 	fmt.Println("-------------if and switch statements---------")
// 	pToggle := true
// 	if pToggle {
// 		fmt.Println("Flag pToggle is true.")
// 	}
// 	statePopulations := map[string]int{
// 		"Maharashtra": 999999,
// 		"Goa":         3333,
// 		"Gujarat":     4232435,
// 	}

// 	// The semicolon in the if statement first serves as the initializer
// 	if population, stateExists := statePopulations["Karnataka"]; stateExists {
// 		fmt.Println(population)
// 	} else {
// 		fmt.Println("State does not exist in collection.")
// 	}

// 	// The scope of population is only valid in that if statment.
// 	// So below code print statement would give an error.
// 	//fmt.Println(population)
// }

// Guess the number game -- hardcoded.
// func main() {
// 	fmt.Println("-----------Guess the number-----------")
// 	correctNumber := 46
// 	guessNumber := -5
// 	if guessNumber < 1 || guessNumber > 100 {
// 		fmt.Println("The no must be between 1 and 100.")
// 	} else {
// 		if guessNumber < correctNumber {
// 			fmt.Println("Too low...")
// 		}
// 		if guessNumber > correctNumber {
// 			fmt.Println("Too high...")
// 		}
// 		if guessNumber == correctNumber {
// 			fmt.Println("You won, Congratulations...")
// 		}
// 		fmt.Println(guessNumber == correctNumber)
// 	}
// }

// Time stamp : 3:06
// Switch case

// func main() {
// 	fmt.Println("-----------------Switch case-----------------")
// 	isBarryFast := "depends"
// 	switch isBarryFast {
// 	case "no":
// 		fmt.Println("Barry is not that fast.")
// 	case "yes":
// 		fmt.Println("Barry is definitely fast")
// 	default:
// 		fmt.Println("What do you mean by that ?")
// 	}

// }

// Timestamp : 3:08
// Example 2
// func main() {
// 	fmt.Println("---------------Switch Case------------")
// 	// The switch statment cannot have an assignment irrespective
// 	// of whether it requires evaluation. To work with
// 	// this we can add a semicolon and use the variable again.
// 	switch i := 2 + 3; i {
// 	case 1, 5, 10:
// 		fmt.Println("One, five or ten.")
// 	case 2, 6, 9:
// 		fmt.Println("Two, six or nine")
// 	default:
// 		fmt.Println("Another number...")
// 	}

// 	i := 4
// 	switch {
// 	case i <= 10:
// 		fmt.Println("less than or equal to ten.")
// 		fallthrough
// 	case i < 20:
// 		fmt.Println("less than 20.")
// 	case i < 5:
// 		// This statement does not get printed because fallthrough is logicless
// 		// It just gets the next case true. It does not mean that see
// 		// if other cases also match.
// 		// source : https://golang.org/ref/spec#Fallthrough_statements
// 		fmt.Println("less than 5")
// 	default:
// 		fmt.Println("Greater than 20.")
// 	}

// 	var s interface{} = []int{45, 67, 78}
// 	switch s.(type) {
// 	case int:
// 		fmt.Println("s is an integer.")
// 	case float64:
// 		fmt.Println("s is a float64")
// 	case string:
// 		fmt.Println("s is a string.")
// 	// For arrays, their size should match to make the case true.
// 	case []int:
// 		fmt.Println("s is a slice of integers.")
// 	default:
// 		fmt.Println("s is of another type.")
// 	}
// }

// ----------------------- Looping ------------------------
// There is only one statement in golang for looping and that is for loop
// func main() {
// 	fmt.Println("----------------------- Looping ------------------------")
// 	// Simple looping with step as an expression.
// 	// for i := 0; i < 5; i = i + 2 {
// 	// 	fmt.Println(i)
// 	// }
// 	// fmt.Println("----------------------- Example 2 ------------------------")
// 	// Looping when 2 variables are involved.
// 	// Don't use i++ because it is a statement and not an expression.
// 	// Multiple assigned expressions are allowed but syntactically two
// 	// statements like i++,j++ is not allowed.
// 	// for i, j := 0, 0; j < 5; i, j = i+1, j+1 {
// 	// 	fmt.Println(i, j)
// 	// }

// 	// Expressing a while loop with condition:
// 	// i := 0
// 	// for i < 5 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }
// 	// break and continue can be used in the same way it is used in "C"
// 	// outerloop:
// 	// 	for i := 1; i <= 3; i++ {
// 	// 		for j := 1; j <= 3; j++ {
// 	// 			fmt.Println(i * j)
// 	// 			if i*j >= 3 {
// 	// 				// Only the second loop will break, the first one will still be active.
// 	// 				// break
// 	// 				// Use a label for the outermost loop and use it in the break statement.
// 	// 				break outerloop
// 	// 			}
// 	// 		}
// 	// 	}
// 	fmt.Println("----------------------- Looping on collections ------------------------")
// 	s := []int{1, 2, 3}
// 	fmt.Println(s)
// 	for k, v := range s {
// 		fmt.Println(k, v)
// 	}
// 	sample := "Yo, what's up?"
// 	// Use the _ operator to get only the value.
// 	for _, v := range sample {
// 		fmt.Println(string(v))
// 	}

// 	statePopulations := map[string]int{
// 		"Maharashtra": 999999,
// 		"Goa":         3333,
// 		"Gujarat":     4232435,
// 	}
// 	// Printing only keys is allowed in Go.
// 	for k := range statePopulations {
// 		fmt.Println(k)
// 	}
// }

// func main() {
// 	fmt.Println("------------------ Defer ------------------------")
// 	// Below statements are printed in reverse order.
// 	// the defer statement works differently.
// 	// Each time a "defer" statement executes, the function value and parameters
// 	// to the call are evaluated as usual and saved anew but the actual function is
// 	// not invoked. Instead, deferred functions are invoked immediately before the
// 	// surrounding function returns, in the reverse order they were deferred.
// 	// https://golang.org/ref/spec#Defer_statements
// 	defer fmt.Println("start")
// 	defer fmt.Println("middle")
// 	defer fmt.Println("end")
// }

// func main() {
// 	fmt.Println("------------------ Defer ------------------------")
// 	res, err := http.Get("https://www.google.co.in/robots.txt")
// 	if err != nil {
// 		fmt.Println("Error detected...")
// 		log.Fatal(err)
// 	}
// 	// If you want the resource to be open because
// 	// you may want to process something the string in this case
// 	// to filter out the stuff you need. We can use defer statement
// 	// here to close the resource whenver you are done with all
// 	// processing automatically.
// 	defer res.Body.Close()
// 	robots, err := ioutil.ReadAll(res.Body)
// 	// res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }

// func main() {
// 	fmt.Println("------------------ Defer ------------------------")
// 	a := "start"
// 	defer fmt.Println(a)
// 	a = "end"
// 	defer fmt.Println(a)
// 	// you think "a" will result to "end" because the by the time
// 	// but it doesn't work that way. defer evaulates the function
// 	// but invokes it only before the surrounding function is going to exit.
// }

// func main() {
// 	fmt.Println(".............Panic and recover..........")
// 	// a, b := 1, 0
// 	// c := a / b
// 	// //panic("Something doesn't seem right...")
// 	// fmt.Println(c)

// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.Write([]byte("Hello there..."))
// 	// })
// 	// // If we dont use the panic statement and run the code
// 	// //http.ListenAndServe(":8080", nil)
// 	// err := http.ListenAndServe(":8080", nil)
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	fmt.Println("start")
// 	panicker()
// 	fmt.Println("end")
// }

// func panicker() {
// 	fmt.Println("About to panic")
// 	// Anonymous function.
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println("Error : ", err)
// 			// Rethrow the panic if you know you cannot handle this error at runtime.
// 			// panic(err)
// 		}
// 	}()
// 	// Above function which is deferred will be executed just before panicking
// 	panic("Something went wrong here...")
// 	// Execution will never reach below
// 	fmt.Println("Done panicking.")
// }

// Pointers
// func main() {
// 	fmt.Println("------------Pointers--------------")
// 	// var a int = 42
// 	// // b is not holding any value, it points to the memory address of a.
// 	// var b *int = &a
// 	// fmt.Println(a, b)
// 	// fmt.Println(a, *b)
// 	// *b = 10
// 	// fmt.Println(a, *b)

// 	// a := [3]int{1, 2, 3}
// 	// b := &a[0]
// 	// d := &a[2]
// 	// c := &a[1]
// 	// // Go does not allow pointer arithmatic
// 	// // c := &a[1] - 4
// 	// fmt.Println(a, *b, *c, *d)
// 	// fmt.Printf("%T %T %T %T\n", a, b, c, d)
// 	// fmt.Println(*c)
// 	// fmt.Println(*d)

// 	a := []int{1, 2, 3}
// 	b := a
// 	fmt.Println(a, b)
// 	a[1] = 56
// 	// A slice doesn't hold the actual value of the array unlike arrays which hold values
// 	// and the size. A slices is a pointer to the first element of the sequence.
// 	fmt.Println(a, b)
// 	// Similar behaviour is observed in maps.
// }

// Functions
// Timestamp : 04:21:00
// func main() {
// 	// fmt.Println("--------------- Functions ------------------")
// 	// fmtMessage("Anagha", "What's up?", 4)

// 	// greeting := "Hello"
// 	// name := "Anagha"
// 	// sayGreeting(&greeting, &name)
// 	// fmt.Println(name)

// 	// Passing Variadic parameters
// 	// In case of multiple arguments, keep the variadic parameters at the end.
// 	// s := sum("The result is : ", 1, 2, 3, 4, 5)
// 	// fmt.Println(s)
// 	// Passing pointers for the same.
// 	// s := sum("The result is : ", 1, 2, 3, 4, 5)
// 	// fmt.Println("The result is :", *s)

// 	// Example for multiple return types.
// 	// d, err := divide(1, 0)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(d)

// 	// Example for anonymous functions
// 	// for i := 0; i < 5; i++ {
// 	// 	// It is standard practice to pass i as the parameter so that
// 	// 	// it behaves properly as expected in case the function is executed
// 	// 	// in asynchronous mode.
// 	// 	func(i int) {
// 	// 		fmt.Println(i)
// 	// 	}(i)
// 	// }

// 	// var divide func(float64, float64) (float64, error)
// 	// divide = func(a, b float64) (float64, error) {
// 	// 	if b == 0.0 {
// 	// 		return 0.0, fmt.Errorf("Cannot divide by zero.")
// 	// 	} else {
// 	// 		return a / b, nil
// 	// 	}
// 	// }
// 	// d, err := divide(5.0, 0.0)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(d)

// 	g := greeter{
// 		greeting: "Hello",
// 		name:     "Sarah",
// 	}
// 	g.greet()
// 	fmt.Println("The new name is :", g.name)

// }

// We can define arguments with the same datatype only once in a consecutive order
// the compiler will infer the individual data type.
func fmtMessage(name, msg string, assignmts_nos int) {
	fmt.Println("-----------------------------------")
	fmt.Println("***********************************")
	fmt.Println("Hey " + name + "! " + msg)
	fmt.Println("Are you done with those " + fmt.Sprint(assignmts_nos) + " assignments ?")
	fmt.Println("***********************************")
	fmt.Println("-----------------------------------")
}

// Passing pointers directly can save memory and also make your program more efficient.
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Madhura"
	fmt.Println(*name)
}

// Variadic parameters.
// The int after the function definition is required for setting the data type of the return variable
func sum(msg string, values ...int) *int {
	// fmt.Println(values)
	// fmt.Printf("%T\n", values)
	result := 0
	for _, v := range values {
		result += v
	}
	// fmt.Println(msg, result)
	// return result

	// In other languages the local stack of a function is destroyed when the function is done
	// executing. Golang saves this on a shared memory(heap memory) before destroying the local stack
	return &result
}

// Multiple return parameters.
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Denominator cannot be zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// Below statement is a method not a function.
// A function is called a method when it operates within a defined context.
// This method below will print out elements inside the struct if assigned.
func (res *greeter) greet() {
	fmt.Println(res.greeting, res.name)
	// We can also pass the reference for changing the actual value itself
	// rather than it's copy.
	res.name = "Anagha"
}

func checkMode(mode string) {
	for i := 0; i < 3; i++ {
		fmt.Println(mode, " :: ", i)
	}
}

func sayHello() {
	fmt.Println("Hello there : # ", counter)
	m.RUnlock()
	wg.Done()
}

func incrementCounter() {
	counter++
	m.Unlock()
	wg.Done()
}

var wg = sync.WaitGroup{}
var counter int = 0
var m = sync.RWMutex{}

// ------------------ Goroutines -------------------
// func main() {
// 	fmt.Println("--------------- Goroutines ---------------")
// 	// mode := "direct"
// 	// checkMode(mode)

// 	// go checkMode("goroutine")

// 	// go func(msg string) {
// 	// 	fmt.Println(msg)
// 	// }("inline_func_call")

// 	// go checkMode(mode)
// 	// mode = "isaltered"
// 	// time.Sleep(10 * time.Millisecond)
// 	// fmt.Println("Done...")

// 	// The go routine prepares the execution thread but doesnt give any attention to it
// 	// immediately. The main function still keeps executing. This is the reason why
// 	// the variable gets a new value while the function is executing the print statement.
// 	msg := "Hello"
// 	wg.Add(2)
// 	go func() {
// 		fmt.Println(msg, "parameter not passed")
// 		// Done decreaments the wait group counter by 1.
// 		wg.Done()
// 	}()
// 	// But when you pass a parameter to the same function and use a goroutine.
// 	// It takes the current value passed but does not execute it.
// 	go func(message string) {
// 		fmt.Println(message, "parameter passed")
// 		wg.Done()
// 	}(msg)
// 	msg = "world"

// 	// time.Sleep(100 * time.Millisecond)

// 	// In real world scenarios, we cannot use a forceful delay statement. In those cases
// 	// we can use waitgroups.
// 	wg.Wait()
// }

// This below example has multiple race conditions and the result of this program
// will always be different every time we run it.
func main() {
	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go sayHello()
	// 	go incrementCounter()
	// }
	// wg.Wait()

	// There is a way to handle the race conditions by introducing mutex.
	// mutex object allows all the processes to use the same resource but
	// at a time only one process is allowed to use the resource.
	for i := 0; i < 100; i++ {
		runtime.GOMAXPROCS(8)
		wg.Add(2)
		m.RLock()
		fmt.Println("go sayHello...")
		go sayHello()
		m.Lock()
		fmt.Println("go incrementcounter...")
		go incrementCounter()
	}
	wg.Wait()
}
