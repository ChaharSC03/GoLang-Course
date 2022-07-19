package main

import "fmt"

func main() {

	defer fmt.Println("thanks 1 ")
	defer fmt.Println("thank you 2 ")
	// defer works as last in first out
	// the last define defer will execute first in the programm

	fmt.Println("welcome to defers in golangs")
	def()

}

func def() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		// fmt.Println(i)
	}
}
