package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Golang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our Chai:")

	// comma ok syntax || err ok
	// no try catch in golang
	// _ is used to ignore the return value
	// inputRating, err := reader.ReadString('\n')

	inputRating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", inputRating)
	fmt.Printf("Type of imputRating is, %T", inputRating)
}
