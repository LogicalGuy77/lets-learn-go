package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Cafe!")
	fmt.Println("Enter the rating for our Cafe: [0-5]")
	reader := bufio.NewReader(os.Stdin)
	inputRating, _ := reader.ReadString('\n')
	// we are getting error because of the new line character at the end of the input
	fmt.Println("Thanks for rating, ", inputRating)

	// Conversion
	newRating, err := strconv.ParseInt(strings.TrimSpace(inputRating), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("New rating is %d, type of newRating is %T", newRating+1, newRating)
	}
}
