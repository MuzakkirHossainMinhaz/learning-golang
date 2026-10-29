package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to if-else in Go")

	fmt.Println("\n1. Basic if, else if, and else")
	loginCount := 10
	if loginCount < 10 {
		fmt.Println("Regular user")
	} else if loginCount > 10 {
		fmt.Println("Watch out")
	} else {
		fmt.Println("Exactly 10 login attempts")
	}

	fmt.Println("\n2. Conditions must be bool expressions")
	score := 82
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: below C")
	}

	fmt.Println("\n3. Combine conditions with logical operators")
	age := 20
	hasID := true
	if age >= 18 && hasID {
		fmt.Println("Entry allowed")
	} else {
		fmt.Println("Entry denied")
	}

	if age < 13 || age >= 65 {
		fmt.Println("Discount eligible")
	} else {
		fmt.Println("Standard price")
	}

	fmt.Println("\n4. Short-circuit evaluation")
	var name *string
	// The second expression is evaluated only if name != nil.
	if name != nil && *name != "" {
		fmt.Println("Name:", *name)
	} else {
		fmt.Println("Name is missing")
	}

	fmt.Println("\n5. An if initializer keeps a variable's scope local")
	if number := 3; number < 10 {
		fmt.Printf("%d is less than 10\n", number)
	} else {
		fmt.Printf("%d is not less than 10\n", number)
	}
	// number is not available here.

	fmt.Println("\n6. Use if with the comma-ok map lookup idiom")
	scores := map[string]int{"Go": 95}
	if score, found := scores["Go"]; found {
		fmt.Printf("Go score: %d\n", score)
	} else {
		fmt.Println("No score found")
	}

	fmt.Println("\n7. Handle errors with if")
	if value, err := strconv.Atoi("42"); err != nil {
		fmt.Println("Could not parse number:", err)
	} else {
		fmt.Printf("Parsed number: %d\n", value)
	}

	fmt.Println("\n8. Guard clauses return early on invalid input")
	if result, err := divide(10, 2); err != nil {
		fmt.Println("Division failed:", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}
	if _, err := divide(10, 0); err != nil {
		fmt.Println("Division failed:", err)
	}

	fmt.Println("\n9. Use if/else to choose a value")
	number := 7
	parity := "odd"
	if number%2 == 0 {
		parity = "even"
	}
	fmt.Printf("%d is %s\n", number, parity)
}

func divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("denominator cannot be zero")
	}

	return numerator / denominator, nil
}
