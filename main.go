package main 

import "fmt"

// Cannot declare nameFive := "Cannot be declared here"

// Only this format can be declared outside
var nameFive = "Can be declared here"

// Varibales
// Integers
func main() {
	
	// Strings
	var nameOne string = "hello"
	var nameTwo = "Sammy"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "How's you?"
	nameThree = "Beloved"

	fmt.Println(nameOne, nameTwo, nameThree)

	// := cannot be used outside of a function
	nameFour := "Wow, awesome!"

	fmt.Println(nameFour)

	// ints
	var ageOne int = 22
	var ageTwo = 25
	ageThree := 30

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 15
	var numTwo int8 = -128

	// When using uint - cannot use negative integers
	var numThree uint64 = 12423423232324234234

	// float
	var scoreOne float32 = 25.89
	var scoreTwo float64 = 4343493434343453.6
	scoreThree := 1.4

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// Print Statements
	
	// Print
	// This Print format does not include new lines
	fmt.Print("Hello, ")
	fmt.Print("World!")
	// It will print "Hello, World!"
	// To insert a new line we have to add the \n for new line
	fmt.Print("How is the world \n")
	// Output
	// "Hello, World!"
	// "How is the world"

	age := 22
	name := "Sammy"

	// Println
	// Automatically inserts a new line
	fmt.Println("Hello, guys")
	fmt.Println("Goodbye, guys")
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	// Does not add new lines automatically
	// %v is for variable
	// %q is quote
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
}