package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https//lco.dev:3000/learn?coursename=reactjs&payid=12afas21fa"

func main() {
	fmt.Println("handling URL in golang")

	// fmt.Println(myurl)

	// parsing --> extract some info or stuff

	result, err := url.Parse(myurl)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qpara := result.Query()
	fmt.Printf("type of qpara %T \n", qpara)

	fmt.Println(qpara["payid"])

	for _, val := range qpara {
		fmt.Println("qpara is:", val)
	}

	partofURL := &url.URL{
		Scheme: "https",
		Host:   "instagram.com",
		Path:   "/vivek_prasad_03",
	}

	anotherURl := partofURL.String()

	fmt.Println(anotherURl)

}
