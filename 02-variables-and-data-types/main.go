package main

import "fmt"

const LoginToken string = "constant_token_login"

func main() {
	var username string = "minhaz"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username) // %T is used to print the type of the variable

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallIntValue uint8 = 255
	fmt.Println(smallIntValue)
	fmt.Printf("Variable is of type: %T \n", smallIntValue)

	var smallFloatValue float64 = 255.45544511254451885
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n", smallFloatValue)

	// default values and some aliases
	var variable int
	fmt.Println(variable)
	fmt.Printf("Variable is of type: %T \n", variable)

	// implicit type
	var fullname = "Muzakkir Hossain Minhaz"
	fmt.Println(fullname)

	// no var style
	balanceOfUser := 300000.0
	fmt.Println(balanceOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
