package main

import "fmt"

// this will not work outside of a function
// jwtToken := sbdjhbsjdbj

// Public variable
const LoginToken string = "gibrissudbi"

func main() {
	fmt.Println("Variables")

	var username string = "Meet"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T \n", smallInt)

	var smallFloat float32 = 35.35164641645649861653
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)

	// implicit type
	var website = "youtube.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	
	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)
	
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
