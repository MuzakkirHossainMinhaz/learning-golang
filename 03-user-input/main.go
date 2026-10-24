package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Pizza Rating App"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your pizza rating (1-5):")

	// comma-ok || comma-err syntax

	// ReadString will read until the delimiter is found, in this case '\n' (newline character)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating")
	fmt.Printf("You rated the pizza %s", input) // include the newline character in the output
	fmt.Printf("Type of the input is %T", input)
}
