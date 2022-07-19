package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ChaharSC03/mongodbapi/router"
)

func main() {
	fmt.Println("mongoDB api")

	fmt.Println("server is ready to start--")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":8000", r))

	fmt.Println("listenig to the port:8000")
}
