package main

import "fmt"

// In golang public variables/func start with capital letter and private with small letter
const Pi float32 = 3.14

func main() {
	var username string = "Harshit"
	fmt.Println("My name is", username)
	fmt.Printf("Data type of username is %T\n", username)

	// default values and some aliases
	var age int
	fmt.Println("My age is", age)
	fmt.Printf("Data type of age is %T\n", age)

	var address string
	fmt.Println("My address is", address)
	fmt.Printf("Data type of address is %T\n", address)

	// implicit type
	var country = "India"
	fmt.Println("I am from", country)

	// no var style := this is not allowed outside a function
	// := walrus operator
	numberOfBlocs := 25
	fmt.Println("Number of blocks are", numberOfBlocs)

	fmt.Println("Value of Pi is", Pi)
	fmt.Printf("Data type of Pi is %T\n", Pi)
}
