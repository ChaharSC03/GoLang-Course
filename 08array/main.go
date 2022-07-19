package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	// array
	var fruitList [4]string

	fruitList[0]="orange"
	fruitList[1]="apple"
	fruitList[2]="guvava"
	fruitList[3]="moango"

	fmt.Println("frunit list is =",fruitList)
	fmt.Println("length of frunit list is =",len(fruitList))


	var veg =[3]string{"tomato", "loki", "potato"}
	fmt.Println("vegetable list is",veg, "\nand length is",len(veg))
}
