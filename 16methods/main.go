package main

import "fmt"

func main() {
	fmt.Println("welcome to the methods in golang")
	
	sumit:= User{"sumit chahar",18,12,"aiml",false}
	sumit.GetStatus()
	sumit.getCourse()

}

type User struct {
	Name string
	Age int
	Roll int
	Course string
	Student bool

}

func (u User)  GetStatus(){
	fmt.Println("is user active",u.Student)
}

func (u User)  getCourse(){
	fmt.Println("the course is:",u.Course)
}

