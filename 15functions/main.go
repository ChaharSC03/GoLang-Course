package main

import "fmt"

func main() {
	fmt.Println("welcome to the function in golang")

	// greet()

	sum := add(3, 4)
	fmt.Println("the sum of number is:", sum)

	result := proadder(3, 4, 5)
	fmt.Println("the sum of number is:", result)

}

// to return mention the type after parameter braces of function
func add(a int, b int) int {
	var c int = a + b
	// fmt.Println("sum is =",c)
	return c
}

// if you dont know the no of values for function
func proadder(values ...int) int {
	total := 0
	// here the values as a parameter become the slice so loop through it
	for _, val := range values {
		total = total + val
	}

	return total
}

// func greet()  {
// 	fmt.Println("namaste from India")
// }
