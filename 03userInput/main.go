package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome:="welcome to the userInput"
	fmt.Println(welcome)

	// bufio is the package use to read user input on Standard input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the input for the Program:")
	// comma ok || err ok   - > syntax to read input

	input,_ := reader.ReadString('\n') 
	// input from the reader will be of type string ...

	fmt.Println("the input is:",input)

	// input,err := reader.ReadString('\n')
	// fmt.Println(err)


	getinput :=bufio.NewReader(os.Stdin)

	fmt.Println("enter the input for the program")

	inptt,_ := getinput.ReadString('\n')

	fmt.Println("the input is: ",inptt)






}