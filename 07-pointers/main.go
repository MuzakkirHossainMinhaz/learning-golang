package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) celebrateBirthday() {
	p.age++
}

func addOne(value *int) {
	if value == nil {
		fmt.Println("Cannot update the number: the pointer is nil.")
		return
	}

	*value++
}

func main() {
	fmt.Println("Go Pointers: A Guided Example")

	fmt.Println("\n1. A pointer's zero value is nil")
	var nilPointer *int
	fmt.Printf("Uninitialized pointer: %v\n", nilPointer)
	fmt.Printf("Is it nil? %t\n", nilPointer == nil)

	fmt.Println("\n2. Get an address and read the pointed-to value")
	score := 23
	scorePointer := &score // & gets the memory address of score.
	fmt.Printf("score: %d\n", score)
	fmt.Printf("Address of score: %p\n", scorePointer)
	fmt.Printf("Value at that address: %d\n", *scorePointer) // * follows the address.

	fmt.Println("\n3. Change a value through its pointer")
	*scorePointer += 2
	fmt.Printf("score after updating through the pointer: %d\n", score)

	fmt.Println("\n4. Copying a pointer still refers to the same value")
	anotherScorePointer := scorePointer
	*anotherScorePointer = 30
	fmt.Printf("score: %d\n", score)
	fmt.Printf("Value through the original pointer: %d\n", *scorePointer)
	fmt.Printf("Value through the copied pointer: %d\n", *anotherScorePointer)

	fmt.Println("\n5. Reassigning a pointer does not change the original variable")
	anotherScorePointer = new(int)
	*anotherScorePointer = 7
	fmt.Printf("score is still: %d\n", score)
	fmt.Printf("The reassigned pointer now points to: %d\n", *anotherScorePointer)

	fmt.Println("\n6. Pass a pointer to a function to update the original value")
	addOne(&score)
	fmt.Printf("score after addOne: %d\n", score)
	addOne(nilPointer) // A nil check prevents dereferencing a nil pointer.

	fmt.Println("\n7. Create a pointer with new")
	counterPointer := new(int) // new allocates an int initialized to its zero value.
	fmt.Printf("Counter before update: %d\n", *counterPointer)
	*counterPointer = 10
	fmt.Printf("Counter after update: %d\n", *counterPointer)

	fmt.Println("\n8. Pointers to structs and pointer receiver methods")
	alex := person{name: "Alex", age: 30}
	personPointer := &alex
	fmt.Printf("Person before update: %+v\n", alex)
	personPointer.age++ // Go allows field access through a struct pointer.
	personPointer.celebrateBirthday()
	fmt.Printf("Person after two pointer-based updates: %+v\n", alex)

	fmt.Println("\n9. Take the address of a struct literal")
	samPointer := &person{name: "Sam", age: 25}
	fmt.Printf("New person: %+v\n", *samPointer)

	fmt.Println("\n10. Two pointers can refer to the same value")
	aliasPointer := &score
	*aliasPointer = 99
	fmt.Printf("score: %d; value through scorePointer: %d\n", score, *scorePointer)

	fmt.Println("\n11. A pointer can point to another pointer")
	pointerToScorePointer := &scorePointer
	fmt.Printf("Address stored in pointerToScorePointer: %p\n", pointerToScorePointer)
	fmt.Printf("Value reached through two pointers: %d\n", **pointerToScorePointer)

	fmt.Println("\n12. Go pointers do not support pointer arithmetic")
	fmt.Println("Use indexes and slices to move through collections instead.")
}
