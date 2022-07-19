package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("conversion")
	// user input

	getInput := bufio.NewReader(os.Stdin) //refernce for reader

	fmt.Println("enter the input here ->")

	inputt, _ := getInput.ReadString('\n')
	//input values until user hit new line or enter
	// _ after the input is for error

	fmt.Println("the input is -> ", inputt)

	// package to convert is strconv

	inputts, err := strconv.ParseFloat(strings.TrimSpace(inputt), 64)
	// using package stings use here to trim space from the inputt string

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("adding some number to it after conversion to float ", inputts+4)
	}

}
