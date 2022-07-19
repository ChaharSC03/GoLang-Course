package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)




func main() {
	fmt.Printf("how to make POST request with JSON data in golang ")
	fmt.Println("")

	// PerformPOSTRequest_json()

	PerformPOSTformRequest()
}

func PerformPOSTRequest_json()  {

	const myurl="http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
	
	{
		"course":"golangg",
		"duration": 3,
		"platform": "learncodeonline.in"
	}
	`)
	
	// send request

	response ,err :=http.Post(myurl, "application/json", requestBody)

	if err!=nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPOSTformRequest()  {

	const myurl="http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstName","Sumit")
	data.Add("lastName","Chahar")
	data.Add("Email","Sumit030@gma.com")

	response , err := http.PostForm(myurl,data)

	if err!=nil {
		panic(err)
	}

	defer response.Body.Close()
	
	content ,_ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	
}

