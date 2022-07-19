package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to Working with file in golang")

	cont :="this is golang my friend and we are learning to work with the file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// else {
	// 	io.WriteString(file,cont)
	// }

	len, err :=io.WriteString(file,cont)

	checkerr(err)
	fmt.Println("the length is",len)
	defer file.Close()

	readfile("./myfile.txt")
}

func readfile(filename string)  {
	
	data,err := ioutil.ReadFile(filename)

	checkerr(err)

	fmt.Println("data inside the file is \n",string(data))
}

func checkerr(err error)  {

	if err!=nil {
		panic(err)
	}
}
