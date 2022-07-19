package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices in golang")

	// slices

	var fruitList=[]string{"apple","mango","orange"}

	fmt.Println("fruitList is :",fruitList)
	fmt.Println("length fof fruitList is :",len(fruitList))
	
	// to add or append the item in the slice list
	fruitList=append(fruitList, "litchi","guava")

	fmt.Println(fruitList)

	// to make seperate part of your slice
	// fruitL :=append(fruitList[2:])
	// fruitList = append(fruitList[2:4])
	// fmt.Println(fruitList)


	// using make() to create slices

	scores := make([]int , 3)
	scores[0] =12
	scores[1] =22
	scores[2] =2

	fmt.Println(scores)

	// to append 
	scores=append(scores, 35,232,11)

	fmt.Println(scores)

	// to check whether it is true or not
	fmt.Println("is this scrore sorted?=",sort.IntsAreSorted(scores))
	// sort package to sort any slice of int string etc

	sort.Ints(scores)
	fmt.Println("the sorted array of scores ")

	fmt.Println(scores)

	fmt.Println("is this scrore sorted?=",sort.IntsAreSorted(scores))



	// remove element from the slices using index

	var course=[]string{"java","c","javascript","go"}

	fmt.Println("\ncourses are : ",course)

	var index = 2

	course = append(course[:index],course[index+1:]...)

	fmt.Println(course)






}