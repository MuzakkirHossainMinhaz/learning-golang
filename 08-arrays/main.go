package main

import "fmt"

// arrayBasics demonstrates how arrays are declared and initialized in Go.
// Arrays have a fixed length, and every element has the same type.
func arrayBasics() {
	fmt.Println("=== 1. Array Basics ===")

	// Fixed-size array of 4 strings.
	fruitList := [4]string{"Apple", "Mango", "Orange", "Peach"}
	fmt.Println("Fruit list:", fruitList)
	fmt.Println("Length:", len(fruitList))
	fmt.Println("Capacity:", cap(fruitList)) // For arrays, len and cap are the same.

	// Array with default zero values for unspecified positions.
	var numbers [5]int
	numbers[0] = 10
	numbers[2] = 20
	fmt.Println("Numbers:", numbers)
}

// arrayIndexing shows how to access and modify elements by index.
func arrayIndexing() {
	fmt.Println("\n=== 2. Indexing and Updating Values ===")

	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println("First color:", colors[0])

	colors[1] = "Yellow"
	fmt.Println("Updated colors:", colors)
}

// arrayIteration shows different ways to loop over arrays.
func arrayIteration() {
	fmt.Println("\n=== 3. Iterating Arrays ===")

	marks := [4]int{88, 76, 90, 81}

	// Standard for loop with index.
	fmt.Println("Using index loop:")
	for i := 0; i < len(marks); i++ {
		fmt.Printf("marks[%d] = %d\n", i, marks[i])
	}

	// range loop gives index and value.
	fmt.Println("Using range loop:")
	for index, value := range marks {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

// arrayAsFunctionParameter demonstrates passing arrays to functions.
// Arrays are copied when passed by value, so changes inside the function do not affect the original.
func arrayAsFunctionParameter(arr [3]int) {
	arr[0] = 999
	fmt.Println("Inside function (copied array):", arr)
}

// arrayByReference demonstrates pointer-based update to modify the original array.
func arrayByReference(arr *[3]int) {
	arr[0] = 111
	fmt.Println("Inside function (modified by reference):", *arr)
}

// arrayMatrix shows a multidimensional array in Go.
func arrayMatrix() {
	fmt.Println("\n=== 4. Multi-Dimensional Arrays ===")

	// 2D array: 2 rows, 3 columns.
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Matrix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

// arrayOfStructs shows arrays can contain custom types as well.
func arrayOfStructs() {
	fmt.Println("\n=== 5. Arrays of Structs ===")

	type Student struct {
		Name  string
		Grade int
	}

	students := [3]Student{
		{Name: "Alice", Grade: 90},
		{Name: "Bob", Grade: 85},
		{Name: "Charlie", Grade: 95},
	}

	for _, student := range students {
		fmt.Printf("%s scored %d\n", student.Name, student.Grade)
	}
}

// arrayComparison demo: arrays can be compared only if all elements are comparable.
func arrayComparison() {
	fmt.Println("\n=== 6. Comparing Arrays ===")

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	fmt.Println("a == b:", a == b)
	fmt.Println("a == c:", a == c)
}

// arrayZeroValues demonstrates default values for array elements.
func arrayZeroValues() {
	fmt.Println("\n=== 7. Zero Values ===")

	var empty [4]int
	fmt.Println("Default values:", empty)
	fmt.Println("Zero value for int:", empty[0])
}

func main() {
	fmt.Println("Welcome to Arrays in Go!")

	arrayBasics()
	arrayIndexing()
	arrayIteration()

	values := [3]int{10, 20, 30}
	fmt.Println("\nBefore passing to function:", values)
	arrayAsFunctionParameter(values)
	fmt.Println("After value pass:", values)

	arrayByReference(&values)
	fmt.Println("After reference update:", values)

	arrayMatrix()
	arrayOfStructs()
	arrayComparison()
	arrayZeroValues()

	fmt.Println("\nArrays in Go are fixed-size and efficient, but slices are often preferred when the length may change.")
}
