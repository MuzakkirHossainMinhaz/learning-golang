package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pizza Rating App!")
	fmt.Println("Please rate our pizza from 1 to 5:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	ratingText := strings.TrimSpace(input)             // Remove any leading/trailing whitespace
	numRating, _ := strconv.ParseFloat(ratingText, 64) // Convert the string input to a float64

	fmt.Printf("Thanks for rating our pizza %.1f out of 5!\n", numRating)
	fmt.Printf("Your rating plus 1 is %.1f.\n", numRating+1)
}
