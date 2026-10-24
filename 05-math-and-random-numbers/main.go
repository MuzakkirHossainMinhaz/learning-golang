package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
)

func main() {
	fmt.Println("Welcome to Math and Random Numbers!")

	firstNumber := 2
	secondNumber := 4.5
	convertedNumber := int(secondNumber)

	fmt.Printf("The sum of %d and %.1f (converted to %d) is %d.\n",
		firstNumber, secondNumber, convertedNumber, firstNumber+convertedNumber)

	pseudoRandomNumber := mathrand.Intn(5) + 1
	fmt.Printf("Your pseudo-random number from 1 to 5 is %d.\n", pseudoRandomNumber)

	secureRandomNumber, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(5))
	secureRandomNumber.Add(secureRandomNumber, big.NewInt(1))
	fmt.Printf("Your cryptographically secure random number from 1 to 5 is %s.\n",
		secureRandomNumber.String())
}
