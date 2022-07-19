package main

import "fmt"

const myToken string = "abcdefghijklmnopqrstuvwxyz"

func main() {
	// fmt.Println("variables")

	var username string = "sumit" //string values
	fmt.Println(username)

	// to know the type
	fmt.Printf("the type of the variables is= %T ", username)

	var isLogged bool = true //boolean values
	fmt.Println("\nthe user is logged in? ", isLogged)

	var rollnumber uint8 = 255 //integer values
	fmt.Println("the rollnumber is:", rollnumber)

	var decimalVal float32 = 25.555555555
	fmt.Println("the decimal value is:", decimalVal)

	// default values and some aliases
	var valuess int
	fmt.Println(valuess) // to assign of garbage value

	// implicit type

	var charracters = "valuesssss"
	fmt.Println(charracters)

	// no var style

	stingis := "hello"
	fmt.Println(stingis)

	numebris := 303003
	fmt.Println(numebris)


	fmt.Println(myToken)
}
