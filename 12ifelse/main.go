package main

import "fmt"

func main() {
	fmt.Println("welcome to flow statement IF ELSE")

	loginCount := 1
	var result string
	if loginCount < 2 {
		result = "normal user"	
	} else if loginCount < 10 {
		result = "frequesnt user"
	} else {
		result=" no data"
	}

	fmt.Println(result)


	if 8%2 ==0 {
		fmt.Println("even")
	}else {
		fmt.Println("odd")
	}


	// mostly used
	// if _initialize ; _condition {

	// }

	if num:=6; num <5 {
		fmt.Println("True")
	}else {
		fmt.Println("false")
	}
}