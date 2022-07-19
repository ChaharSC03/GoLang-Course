package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("welcome to the maps in golang")

	state := make(map[string]string)

	state["UP"]="Uttar Pradesh"
	state["AP"]="Andhra Pradesh"
	state["MH"]="Maharashtra"

	fmt.Println("list of states in maps:")
	fmt.Println(state)
	fmt.Println("accessing one element",state["UP"])

	// to delete value
	delete(state,"MH")

	fmt.Print(state,"\n")

	// looping in the maps
	// loop
	for key , values := range state{
		fmt.Printf("for key %v , value is %v \n",key,values)
	}

}
