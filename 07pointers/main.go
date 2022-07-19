package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers in golang")

	// var first *int 
	// fmt.Println("value of the pointer: ",first)

	mynumber := 32

	var ptr = &mynumber //providing the reference to the ptr

	fmt.Println("value of the pointer:",*ptr) //pointing to the value
	fmt.Println("address of the pointer:",ptr)


	*ptr= *ptr +2
	fmt.Println("operation on the pointer=",mynumber)


}
