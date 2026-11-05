package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer in Go")

	fmt.Println("\n1. Deferred calls run when their function returns")
	showDeferOrder()

	fmt.Println("\n2. Deferred function arguments are evaluated immediately")
	showArgumentEvaluation()

	fmt.Println("\n3. A deferred closure reads variables when it runs")
	showClosureCapture()

	fmt.Println("\n4. Defers run in last-in, first-out (LIFO) order")
	showDeferredLoop()

	fmt.Println("\n5. A defer runs on an early return")
	showEarlyReturn()

	fmt.Println("\n6. A deferred function can modify named return values")
	fmt.Println("Returned value:", namedReturn())

	fmt.Println("\n7. A deferred function can recover from a panic")
	recoverFromPanic()

	// defer is useful for cleanup after acquiring a resource, for example:
	// file, err := os.Open(path)
	// if err != nil {
	//     return err
	// }
	// defer file.Close()
	//
	// A deferred call does not run if the program exits via os.Exit.
}

func showDeferOrder() {
	defer fmt.Println("First defer runs last")
	defer fmt.Println("Second defer runs first")

	fmt.Println("Function body runs before its deferred calls")
}

func showArgumentEvaluation() {
	message := "value when defer is reached"
	defer fmt.Println("Deferred argument:", message)

	message = "value changed afterward"
	fmt.Println("Current value:", message)
}

func showClosureCapture() {
	message := "value before return"
	defer func() {
		fmt.Println("Deferred closure sees:", message)
	}()

	message = "value at function return"
	fmt.Println("Current value:", message)
}

func showDeferredLoop() {
	for i := 0; i < 4; i++ {
		// Each argument is captured now; the calls run in reverse order
		// when this function returns.
		defer fmt.Print(i)
	}
	fmt.Println("Loop finished; deferred output follows:")
	// The deferred numbers print as 3210 after this line.
}

func showEarlyReturn() {
	defer fmt.Println("Deferred call still runs after an early return")
	fmt.Println("Returning before the function's end")
	return
}

// Named result parameters are variables, so a defer can change them before
// the caller receives the return value.
func namedReturn() (result int) {
	defer func() {
		result++
	}()

	return 5
}

func recoverFromPanic() {
	defer func() {
		// recover only stops a panic when called directly by a deferred function.
		if recovered := recover(); recovered != nil {
			fmt.Println("Recovered panic:", recovered)
		}
	}()

	panic("something went wrong")
}
