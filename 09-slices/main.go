package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Welcome to Slices in Go!")

	sliceBasics()
	sliceLengthAndCapacity()
	slicingAndSharedStorage()
	appendAndCopy()
	sliceIteration()
	values := sliceAsArgument([]int{10, 20, 30})
	fmt.Println("Caller receives:", values)
	sliceUtilities()
	multidimensionalSlices()
}

// A slice is a small descriptor of a contiguous section of an underlying array.
// Unlike an array, a slice's length can change as elements are appended.
func sliceBasics() {
	fmt.Println("\n=== 1. Declaring and Initializing Slices ===")

	// A slice literal does not specify a fixed length.
	fruits := []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Fruits: %v (type %T)\n", fruits, fruits)

	// The zero value of a slice is nil. It has length and capacity zero.
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("nil slice: %v, nil: %t, len: %d, cap: %d\n",
		nilSlice, nilSlice == nil, len(nilSlice), cap(nilSlice))
	fmt.Printf("empty but non-nil slice: %v, nil: %t\n",
		emptySlice, emptySlice == nil)

	// make creates a slice with the requested length and optional capacity.
	scores := make([]int, 3, 5)
	fmt.Printf("Made slice: %v, len: %d, cap: %d\n",
		scores, len(scores), cap(scores))

	// Array elements can be exposed as a slice using a slice expression.
	array := [4]string{"North", "South", "East", "West"}
	directions := array[1:3]
	fmt.Println("Slice of an array:", directions)
}

func sliceLengthAndCapacity() {
	fmt.Println("\n=== 2. Length and Capacity ===")

	values := make([]int, 2, 4)
	values[0], values[1] = 10, 20
	fmt.Printf("Initial: %v, len: %d, cap: %d\n",
		values, len(values), cap(values))

	// Capacity is how far a slice can grow into its underlying array
	// before append needs to allocate a new array.
	values = append(values, 30)
	fmt.Printf("After append: %v, len: %d, cap: %d\n",
		values, len(values), cap(values))
}

func slicingAndSharedStorage() {
	fmt.Println("\n=== 3. Slicing and Shared Storage ===")

	numbers := []int{10, 20, 30, 40, 50}
	part := numbers[1:4] // Includes index 1; excludes index 4.
	fmt.Printf("numbers[1:4]: %v, len: %d, cap: %d\n",
		part, len(part), cap(part))

	// Slices usually share their underlying array. Updating one view can
	// therefore be visible through another view.
	part[0] = 99
	fmt.Println("Original after updating part:", numbers)

	// A full slice expression limits the resulting capacity as well as length.
	limited := numbers[1:4:4]
	fmt.Printf("Full slice expression: %v, len: %d, cap: %d\n",
		limited, len(limited), cap(limited))

	// Clone makes an independent copy when the values should not share storage.
	independent := slices.Clone(part)
	independent[0] = 777
	fmt.Println("Clone changed:", independent)
	fmt.Println("Original remains:", numbers)
}

func appendAndCopy() {
	fmt.Println("\n=== 4. Appending, Copying, and Clearing ===")

	items := []string{"pen", "pencil"}
	items = append(items, "marker", "eraser")
	fmt.Println("After appending values:", items)

	moreItems := []string{"ruler", "notebook"}
	items = append(items, moreItems...) // ... passes each element as an argument.
	fmt.Println("After appending another slice:", items)

	destination := make([]string, len(items))
	copied := copy(destination, items)
	fmt.Printf("Copied %d elements: %v\n", copied, destination)

	// copy returns the number copied, which is the smaller of the two lengths.
	shortDestination := make([]string, 2)
	fmt.Printf("Copied to shorter slice: %d elements\n",
		copy(shortDestination, items))
	fmt.Println("Short destination:", shortDestination)

	// copy also works when source and destination overlap.
	overlap := []int{1, 2, 3, 4}
	copy(overlap[1:], overlap[:3])
	fmt.Println("Overlapping copy:", overlap)

	// clear sets every element in the slice to its element type's zero value.
	clear(destination)
	fmt.Println("After clear:", destination)
}

func sliceIteration() {
	fmt.Println("\n=== 5. Indexing and Iteration ===")

	colors := []string{"red", "green", "blue"}
	fmt.Println("First color:", colors[0])
	colors[1] = "yellow"

	for index, color := range colors {
		fmt.Printf("colors[%d] = %s\n", index, color)
	}

	// Use _ when only the values or only the indexes are needed.
	fmt.Print("Values only:")
	for _, color := range colors {
		fmt.Print(" ", color)
	}
	fmt.Println()
}

// A slice argument is passed by value, but its descriptor points to shared
// backing storage. Element updates are visible to the caller. Appending may
// allocate a different backing array, so the caller should use the returned slice.
func updateFirst(values []int) []int {
	if len(values) == 0 {
		return values
	}
	values[0] = 100
	return append(values, 400)
}

func sliceAsArgument(values []int) []int {
	fmt.Println("\n=== 6. Slices as Function Arguments and Results ===")

	fmt.Println("Before function call:", values)
	values = updateFirst(values)
	fmt.Println("Returned slice:", values)
	return values
}

func sliceUtilities() {
	fmt.Println("\n=== 7. Common Functions in the Slices Package ===")

	values := []int{4, 1, 3, 1, 2}
	fmt.Println("Original:", values)

	// Sort changes the order in place; the slice must contain ordered values.
	slices.Sort(values)
	fmt.Println("Sorted:", values)
	fmt.Println("Is sorted:", slices.IsSorted(values))
	position, found := slices.BinarySearch(values, 3)
	fmt.Printf("Binary search for 3: index %d, found %t\n", position, found)

	fmt.Println("Contains 2:", slices.Contains(values, 2))
	fmt.Println("Index of 1:", slices.Index(values, 1))
	fmt.Println("Minimum:", slices.Min(values))
	fmt.Println("Maximum:", slices.Max(values))

	// Compact removes consecutive duplicates and returns a shorter slice.
	// Sort first if duplicates that are far apart should also be removed.
	// Use its result; it also clears the unused tail of the original storage.
	values = slices.Compact(values)
	fmt.Println("After Compact:", values)

	left := []int{1, 2, 3}
	right := []int{1, 2, 4}
	// Slices cannot be compared with == except against nil; compare elements with slices.Equal.
	fmt.Println("Equal:", slices.Equal(left, right))
	fmt.Println("Compare:", slices.Compare(left, right)) // Negative means left is smaller.
	fmt.Println("Equal to a clone:", slices.Equal(left, slices.Clone(left)))

	// Insert, Delete, and Replace return the resulting slice.
	edited := slices.Insert([]string{"a", "c"}, 1, "b")
	edited = slices.Delete(edited, 0, 1) // Delete the half-open range [0, 1).
	edited = slices.Replace(edited, 1, 2, "d", "e")
	fmt.Println("After Insert, Delete, and Replace:", edited)

	reversed := slices.Clone(values)
	slices.Reverse(reversed)
	fmt.Println("Reversed copy:", reversed)

	// Grow ensures room for additional elements, but does not change the length.
	grown := slices.Grow([]int{1, 2}, 3)
	fmt.Printf("After Grow: %v, len: %d, cap: %d\n",
		grown, len(grown), cap(grown))

	// Clip reduces capacity to length; it does not copy the elements.
	clipped := slices.Clip(grown)
	fmt.Printf("After Clip: %v, len: %d, cap: %d\n",
		clipped, len(clipped), cap(clipped))
}

func multidimensionalSlices() {
	fmt.Println("\n=== 8. Multidimensional Slices ===")

	// A slice of slices is not necessarily rectangular: each row is a separate slice.
	rows := [][]int{
		{1, 2, 3},
		{4, 5},
		{6},
	}

	for rowIndex, row := range rows {
		fmt.Printf("Row %d: %v\n", rowIndex, row)
	}

	// Appending to an inner slice updates that row's slice value.
	rows[1] = append(rows[1], 7)
	fmt.Println("After appending to row 1:", rows)
}
