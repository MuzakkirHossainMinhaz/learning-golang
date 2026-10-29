package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Go")

	fmt.Println("\n1. The three-clause for loop")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	fmt.Println("\n2. A condition-only for loop works like while in other languages")
	count := 1
	for count <= 3 {
		fmt.Println("Count:", count)
		count++
	}

	fmt.Println("\n3. An infinite loop ends with break")
	tries := 0
	for {
		tries++
		fmt.Println("Try:", tries)
		if tries == 2 {
			break
		}
	}

	fmt.Println("\n4. Use continue to skip the rest of the current iteration")
	for number := 1; number <= 5; number++ {
		if number == 3 {
			continue
		}
		fmt.Println("Number:", number)
	}

	fmt.Println("\n5. Range over a slice gives index and value")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for index, day := range days {
		fmt.Printf("days[%d] = %s\n", index, day)
	}

	fmt.Println("\n6. Ignore an unused range result with the blank identifier")
	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println("\n7. Range over arrays and maps")
	temperatures := [3]int{18, 22, 25}
	for index, temperature := range temperatures {
		fmt.Printf("Temperature %d: %d°C\n", index, temperature)
	}

	prices := map[string]int{"tea": 3, "coffee": 5}
	for item, price := range prices {
		// Map iteration order is unspecified and may differ between runs.
		fmt.Printf("%s costs %d\n", item, price)
	}

	fmt.Println("\n8. Range over a string iterates Unicode code points (runes)")
	for byteIndex, character := range "Golang" {
		fmt.Printf("byte index %d: %c\n", byteIndex, character)
	}

	fmt.Println("\n9. Range over an integer iterates from zero up to, but not including, it")
	for index := range 3 {
		fmt.Println("Index:", index)
	}

	fmt.Println("\n10. Range over a channel receives until the channel is closed")
	values := make(chan int, 3)
	values <- 10
	values <- 20
	values <- 30
	close(values)
	for value := range values {
		fmt.Println("Received:", value)
	}

	fmt.Println("\n11. Labeled break exits an outer loop")
outer:
	for row := 1; row <= 3; row++ {
		for column := 1; column <= 3; column++ {
			if row == 2 && column == 2 {
				break outer
			}
			fmt.Printf("(%d, %d) ", row, column)
		}
	}
	fmt.Println("\nExited both loops")

	fmt.Println("\n12. Labeled continue advances the labeled outer loop")
rows:
	for row := 1; row <= 3; row++ {
		for column := 1; column <= 2; column++ {
			if row == 2 && column == 1 {
				continue rows
			}
			fmt.Printf("row %d, column %d\n", row, column)
		}
	}

	fmt.Println("\n13. goto can jump to a label, but structured loop control is usually clearer")
	value := 1
	for value <= 3 {
		if value == 2 {
			goto finished
		}
		fmt.Println("Value:", value)
		value++
	}

finished:
	fmt.Println("Reached the label")
}
