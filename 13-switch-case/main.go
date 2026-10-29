package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch-case in Go")

	fmt.Println("\n1. A switch evaluates its expression once")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("Dice value: %d\n", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("You can open")
	case 2, 3:
		fmt.Println("Move a few spaces")
	case 4, 5:
		fmt.Println("Move forward")
	case 6:
		fmt.Println("Move forward and roll again")
	default:
		// This branch handles values not covered by any case.
		fmt.Println("Unexpected dice value")
	}

	fmt.Println("\n2. A switch can have an initializer")
	switch now := time.Now(); now.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Printf("%s is the weekend\n", now.Weekday())
	default:
		fmt.Printf("%s is a weekday\n", now.Weekday())
	}
	// now is scoped to the switch and cannot be used here.

	fmt.Println("\n3. An expressionless switch is a readable if/else chain")
	temperature := 27
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature < 20:
		fmt.Println("Cool")
	case temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

	fmt.Println("\n4. Cases can match multiple values")
	letter := 'e'
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf("%c is a lowercase vowel\n", letter)
	default:
		fmt.Printf("%c is not a lowercase vowel\n", letter)
	}

	fmt.Println("\n5. Cases run top to bottom; fallthrough runs the next case")
	switch 3 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		fallthrough // Runs the next case body without checking its case expression.
	case 4:
		fmt.Println("also runs after fallthrough")
	default:
		fmt.Println("other")
	}
	// Go does not fall through automatically. fallthrough is valid only as the
	// final statement in a case and continues to the immediately following case.

	fmt.Println("\n6. A type switch selects a branch using an interface's dynamic type")
	printType("hello")
	printType(42)
	printType(true)
	printType(nil)

	fmt.Println("\n7. break exits a switch; continue still advances the surrounding loop")
	for number := 1; number <= 5; number++ {
		switch number {
		case 2:
			fmt.Println("Skip 2")
			continue // Continues the for loop, not just the switch.
		case 4:
			fmt.Println("Break out of the switch at 4")
			break // Exits only this switch; the loop proceeds to its next iteration.
		default:
			fmt.Println("Number:", number)
		}
	}

	fmt.Println("\n8. A function can return directly from a switch")
	fmt.Println("Day type:", dayType(time.Wednesday))
}

func printType(value any) {
	switch typedValue := value.(type) {
	case nil:
		fmt.Println("The value is nil")
	case string:
		fmt.Printf("%q is a string with %d characters\n", typedValue, len(typedValue))
	case int:
		fmt.Printf("%d is an int\n", typedValue)
	case bool:
		fmt.Printf("%t is a bool\n", typedValue)
	default:
		fmt.Printf("Unsupported type: %T\n", typedValue)
	}
}

func dayType(day time.Weekday) string {
	switch day {
	case time.Saturday, time.Sunday:
		return "weekend"
	default:
		return "weekday"
	}
}
