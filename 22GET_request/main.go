package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("How to make GET request in golang")

	PerformGetRequest()
}

func PerformGetRequest()  {
	const myurl="http://localhost:8000/get" 

	response, err := http.Get(myurl)

	if err!=nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("STATUS :",response.Status)
	fmt.Println("Content length",response.ContentLength)

	// presered way is string builder process
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	bytesCount, _ := responseString.Write(content)

	fmt.Println("byte COunt: ",bytesCount)

	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}