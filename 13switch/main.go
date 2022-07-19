package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to the switch in golang")

	rand.Seed(time.Now().UnixNano())
	var num int = rand.Intn(9)
	fmt.Println("the number genrate is",num)
	switch num {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
		fallthrough
	case 6:
		fmt.Println("six")
	default: 
		fmt.Println("sorry unidentified")

	}

}
