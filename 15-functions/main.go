package main

import "fmt"

func main() {
    fmt.Println("Welcome to Functions in Go")

    // A function with no parameters or return values.
    greeter()

    // Functions can accept parameters and return a value.
    result := adder(3, 5)
    fmt.Println("Sum:", result)

    // A variadic parameter accepts zero or more values.
    proResult, message := proAdder(2, 5, 8, 7, 3)
    fmt.Println(message, proResult)

    // A variadic function can also receive a slice with ... .
    numbers := []int{1, 2, 3}
    fmt.Println("Slice sum:", sum(numbers...))

    // Functions can return multiple values, including an error.
    quotient, remainder, err := divide(10, 3)
    if err != nil {
        fmt.Println("Division error:", err)
    } else {
        fmt.Printf("10 / 3 = %d remainder %d\n", quotient, remainder)
    }

    // Named return values are initialized to their zero values.
    fmt.Println("Named return sum:", namedAdder(4, 6))

    // Functions are values: they can be stored in variables and passed
    // to other functions as arguments.
    operation := adder
    fmt.Println("Function value:", operation(6, 7))
    fmt.Println("Callback result:", apply(6, 7, adder))

    // Anonymous functions can be assigned to variables or called directly.
    double := func(value int) int {
        return value * 2
    }
    fmt.Println("Anonymous function:", double(5))
    fmt.Println("Immediately called:", func(value int) int {
        return value + 1
    }(5))

    // A closure remembers and updates the variable from its enclosing scope.
    next := counter()
    fmt.Println("Counter:", next(), next(), next())

    // Recursion: factorial calls itself until it reaches its base case.
    fmt.Println("Factorial:", factorial(5))

    // A deferred call runs when showDefer returns.
    showDefer()
}

func greeter() {
    fmt.Println("Hello from Go")
}

// When parameter types match, the type can be written once.
func adder(valOne, valTwo int) int {
    return valOne + valTwo
}

// The ...int parameter gathers arguments into a slice.
func proAdder(values ...int) (int, string) {
    total := sum(values...)
    return total, "Pro result:"
}

func sum(values ...int) int {
    total := 0
    for _, value := range values {
        total += value
    }
    return total
}

// Go functions can return multiple values. Errors are commonly returned
// as the final result.
func divide(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, fmt.Errorf("denominator cannot be zero")
    }
    return numerator / denominator, numerator % denominator, nil
}

// Named return values can make a result's purpose clearer.
func namedAdder(a, b int) (sum int) {
    sum = a + b
    return
}

// A function can accept another function as a parameter.
func apply(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// counter returns a closure that retains its own count between calls.
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// factorial demonstrates a recursive function for non-negative integers.
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func showDefer() {
    defer fmt.Println("Deferred call runs as the function returns")
    fmt.Println("showDefer is running")
}