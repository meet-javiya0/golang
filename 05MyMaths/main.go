package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths i golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4

	fmt.Println("The sum of these two number is: ", mynumberOne+int(mynumberTwo))

	// random number

	// using math
	// rand.Seed(time.Now().UnixNano())	// this is deprecated
	// rand.New(rand.NewSource(5))
	// fmt.Println(rand.Intn(5) + 1)

	// using crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
