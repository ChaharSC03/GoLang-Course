package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")
	// slice 
	days := []string{"sunday","monday","tuesday","wednesday","thursday"}

	// fmt.Println(days)
	// for loop
	// for d:=0 ; d<len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// fmt.Println("")
	// // using index i in for loop
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	for i,day :=range days{
		fmt.Printf("index is %v and value is %v \n",i,day)
	}

	// while loop as for loop
	fmt.Println("")
	var count int=1
	for count <10 {
		if count ==5 {
			goto smt
			
		}
		if count > 5 {
			break
		}
		fmt.Println("value is", count)
		count++
	}
	smt: fmt.Println("done")





}