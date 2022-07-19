package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `josn:"websiteMode"`
	Password string   `json:"-"` //to hide the password in json
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("how to create json data in golang")
	// create json data
	// encodeJson()

	// how to consume json data in golang

	fmt.Println("how to consume the json data")

	decodeJson()

}

// func encodeJson() {
// 	course := []course{
// 		{"GOlang bootcamp", 100, "ONLINE", "1234abcd", []string{"web-dev", "js"}},
// 		{"c++ bootcamp", 200, "ONLINE", "4312abcd", []string{"full-stack", "js"}},
// 		{"java bootcamp", 300, "ONLINE", "24134abcd", nil},
// 	}

// 	// package the above data in JSON data

// 	finalJSON, err := json.MarshalIndent(course, "", "\t")

// 	if err != nil {
// 		panic(err)

// 	}

// 	fmt.Printf("%s\n", finalJSON)
// }

func decodeJson() {

	jsonData := []byte(`
	
	{
		"coursename": "c++ bootcamp",
		"Price": 200,
		"Platform": "ONLINE",
		"Password": "4312abcd",
		"Tags": [
				"full-stack",
				"js"		]
	}
	`)

	var goCourse course

	check :=json.Valid(jsonData)

	if check {
		fmt.Println("json is valid")
		json.Unmarshal(jsonData,&goCourse)
		fmt.Printf("%#v \n",goCourse)
	} else {
		fmt.Println("json is not valid")
	}

	fmt.Println("")
	// add data to keyvalue pairs

	var myData map[string]interface{}

	json.Unmarshal(jsonData,&myData)
	fmt.Printf("%#v \n",myData)
	fmt.Println("")
	for key,va := range myData {
		fmt.Printf("\n key is %v and value if %v and type of - is : %T",key,va,va)
	}



	
}
