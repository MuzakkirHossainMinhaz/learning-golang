package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Structs in Go")

	fmt.Println("\n1. A struct groups related fields into a named type")
	var zeroValue User
	fmt.Printf("The zero value of User: %+v\n", zeroValue)

	// Keyed literals are clear and remain valid if the struct gains fields.
	user := User{
		Name:   "Minhaz",
		Email:  "minhaz@cse.engr.bd",
		Status: true,
		Age:    16,
		Address: Address{
			City:    "Meherpur",
			Country: "Bangladesh",
		},
	}
	fmt.Printf("User: %+v\n", user) // %+v prints the struct with field names.
	fmt.Printf("Name: %s,\t email: %s\n", user.Name, user.Email)

	fmt.Println("\n2. Positional literals must follow the exact field order")
	point := Point{10, 20}
	fmt.Printf("Point: %+v\n", point)

	fmt.Println("\n3. Structs can contain other structs")
	fmt.Printf("Address: %s, %s\n", user.Address.City, user.Address.Country)

	fmt.Println("\n4. Embedding promotes fields and methods (composition, not inheritance)")
	account := Account{
		ID:   42,
		User: user,
	}
	fmt.Printf("Account ID: %d; promoted user name: %s\n", account.ID, account.Name)

	fmt.Println("\n5. Struct pointers refer to the original struct")
	userPointer := &user
	userPointer.Email = "minhaz@example.com" // Go automatically follows the pointer for field access.
	fmt.Printf("Updated email: %s\n", user.Email)

	fmt.Println("\n6. Methods can use value or pointer receivers")
	fmt.Printf("Display name: %s\n", user.DisplayName()) // A value receiver works with a copy.
	user.HaveBirthday()                                  // A pointer receiver can update the original.
	fmt.Printf("Age after birthday: %d\n", user.Age)

	fmt.Println("\n7. A constructor function can initialize a struct")
	newUser := NewUser("Ahmad", "ahmad@example.com", 25)
	fmt.Printf("New user: %+v\n", newUser)

	fmt.Println("\n8. Struct tags provide metadata to packages such as encoding/json")
	encoded, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %s\n", encoded)

	fmt.Println("\n9. Struct values can be compared when all their fields are comparable")
	userCopy := user
	fmt.Printf("A copied User equals the original: %t\n", userCopy == user)
	// A struct containing a slice, map, or function cannot be compared with ==.

	fmt.Println("\n10. Anonymous structs are useful for short-lived data shapes")
	result := struct {
		Message string
		Count   int
	}{
		Message: "Done",
		Count:   1,
	}
	fmt.Printf("Result: %+v\n", result)
}

// Uppercase field names are exported; lowercase names are unexported outside this package.
type User struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Status  bool    `json:"status"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Point struct {
	X int
	Y int
}

// Embedding makes User's fields and methods available through Account.
type Account struct {
	User
	ID int `json:"id"`
}

// NewUser is a constructor-style function; Go does not have constructor syntax.
func NewUser(name, email string, age int) User {
	return User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}

// DisplayName uses a value receiver because it only reads the User.
func (u User) DisplayName() string {
	return u.Name
}

// HaveBirthday uses a pointer receiver so the change persists on the original User.
func (u *User) HaveBirthday() {
	u.Age++
}
