package main

import (
	"fmt"
	"strings"
)

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

	fmt.Println(numOne, numTwo, numThree)

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
	points := 254.555

	// Println
	// Automatically inserts a new line
	fmt.Println("Hello, guys")
	fmt.Println("Goodbye, guys")
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	/* 
		Does not add new lines automatically
		%v is variable
		%q is quotes = the variables will be in quotes in they are a string  "Sammy" and in bits "\x16" if it an integer
		%T is Type =  used to return the type of a variable if its an integer(25) will return int if string("sammy") will return string
		%f if float = returns the float value
		%0.1f will return float in 1 decimal point if %0.2f will return 2 dp
	 */
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f \n", points)
	fmt.Printf("you scored %0.2f \n", points)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

	// Arrays and Slices
	// Cannot append to arrays but can append in slices.
	// var ages [3]int = [3]int{20, 24, 30}
	// Can also be written this way
	var ages = [3]int{20, 24, 30}

	names := [4]string{
		"Babypluto",
		"Klaus",
		"Slice",
		"Kenny",
	}

	// Change a string of index[i] in names
	names[1] = "BonBon"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood but can be modified)
	var scores = []int{
		100, 50, 65,
	}
	scores[2] = 25

	// Append in slices
	scores = append(scores, 98)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	
	// can also be appended
	rangeOne = append(rangeOne, "Freddy")

	fmt.Println(rangeOne, rangeTwo, rangeThree)


	// STANDARD LIBRARY


}

func greeting() {
	greeting := "Hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
}