package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("welcome to maths, crypto and random number")

	// bad practice
	// var numberOne int = 1
	// var numberTwo float64= 2.6
	// fmt.Println("the sum is:",numberOne+ int(numberTwo))

	// random number using math package
	// rand.Seed(300)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// random number using crypto

	randNumber, _ := rand.Int(rand.Reader, big.NewInt(500))

	fmt.Println("the random number is: ",randNumber)

}
