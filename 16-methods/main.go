package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods in Go")

	user := User{
		Name:   "Minhaz",
		Email:  "minhaz@cse.engr.bd",
		Status: true,
		Age:    26,
	}

	fmt.Println("\n1. A method is a function associated with a receiver type")
	fmt.Printf("User: %+v\n", user)
	fmt.Println("Active:", user.GetStatus())
	fmt.Println("Greeting:", user.Greeting())

	fmt.Println("\n2. A value receiver operates on a copy")
	user.ChangeEmailCopy("copy@example.com")
	fmt.Println("Original email after value-receiver method:", user.Email)

	fmt.Println("\n3. A pointer receiver can modify the original value")
	user.ChangeEmail("updated@example.com")
	user.HaveBirthday()
	fmt.Printf("After pointer-receiver methods: email=%s age=%d\n", user.Email, user.Age)
	// Go takes the address of an addressable value for this method call.
	(&user).HaveBirthday()
	fmt.Println("Age after another birthday:", user.Age)

	fmt.Println("\n4. Methods can belong to any defined type in the same package")
	var temperature Celsius = 23.5
	fmt.Printf("%.1f°C is %.1f°F\n", temperature, temperature.Fahrenheit())

	fmt.Println("\n5. A method value captures a value receiver")
	user.Status = true
	active := user.GetStatus
	user.Status = false
	fmt.Println("Captured status:", active())

	fmt.Println("\n6. A method expression takes the receiver as its first argument")
	greetingMethod := (*User).Greeting
	fmt.Println(greetingMethod(&user))

	fmt.Println("\n7. Types satisfy interfaces by implementing their methods")
	var notifier Notifier = user
	fmt.Println(notifier.Notify())
	// HaveBirthday has a pointer receiver, so *User implements BirthdayCelebrant.
	var birthdayPerson BirthdayCelebrant = &user
	birthdayPerson.HaveBirthday()
	fmt.Println("Age after birthday through interface:", user.Age)

	fmt.Println("\n8. Embedded fields promote their methods")
	account := Account{ID: 42, User: user}
	fmt.Println("Promoted greeting:", account.Greeting())
	fmt.Println("Account notification:", account.Notify())

	fmt.Println("\n9. A method can safely handle a nil pointer receiver")
	var missingUser *User
	fmt.Println(missingUser.Greeting())

	// Go does not support method overloading. Each method name must be unique
	// for its receiver type, and methods cannot be declared on non-local types.
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// GetStatus is a read-only method with a value receiver.
func (u User) GetStatus() bool {
	return u.Status
}

// Greeting uses a pointer receiver so it can also handle a nil *User.
func (u *User) Greeting() string {
	if u == nil {
		return "Hello, anonymous user"
	}
	return "Hello, " + u.Name
}

// ChangeEmailCopy demonstrates that changes to a value receiver affect only
// the receiver's copy, not the original User.
func (u User) ChangeEmailCopy(email string) {
	u.Email = email
	fmt.Println("Email inside value-receiver method:", u.Email)
}

// ChangeEmail uses a pointer receiver to persist a change on the original.
func (u *User) ChangeEmail(email string) {
	u.Email = email
}

func (u *User) HaveBirthday() {
	u.Age++
}

// Celsius is a defined type, so methods can be attached to it.
type Celsius float64

func (c Celsius) Fahrenheit() float64 {
	return float64(c)*9/5 + 32
}

// Notifier is implemented implicitly by any type with a Notify method.
type Notifier interface {
	Notify() string
}

type BirthdayCelebrant interface {
	HaveBirthday()
}

// Notify has a value receiver, so both User and *User implement Notifier.
func (u User) Notify() string {
	return "Notification sent to " + u.Email
}

// Embedding promotes User's methods to Account.
type Account struct {
	User
	ID int
}

func (a Account) Notify() string {
	return fmt.Sprintf("Account %d: %s", a.ID, a.User.Notify())
}
