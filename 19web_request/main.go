package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Handling web request in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("respone of type %T \n",response)
	// fmt.Println(response)

	defer response.Body.Close()
	// call should be close 

	data,err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
