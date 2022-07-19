package main

import "fmt"

func main() {
	fmt.Println("welcome to the structs in golang")

	// no inheritance in golang :)

	sumit :=User{"Sumit",8,"AIML",22}
	fmt.Println(sumit)
	// format 
	fmt.Printf("sumit details are: %+v \n",sumit)
	fmt.Printf("sumit details name is %v \n",sumit.Name)
	fmt.Printf("roll is %v \n",sumit.Rollnumber)
	
	
}

// to define structs

type User struct{
	Name string
	Rollnumber int
	Course string 
	Age int
}