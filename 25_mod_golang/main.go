package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("MOD in golang")

	// gorilla mux is used for production 
	// to install : go get -u github.com/gorilla/mux

	greet()

	r := mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")

	// http.ListenAndServe(":8000",r)

	log.Fatal(http.ListenAndServe(":8000",r))
}

func greet()  {
	fmt.Println("namaste kaise ho")
}

func serveHome(w http.ResponseWriter , r *http.Request)  {
	w.Write([]byte("<h1>welcome to the golang</h1>"))
}



// go mod verify : is to verify the go.sum status for unauthorise changes or update in it

// go list : display the module for the application
// go list all : dump very pkgs for application that are installed 

// go list -m all : display the dependencies of the application on the modules or pkgs

// to know which version of gorilla mux you want to use ----> go list -m -versions github.com/gorilla/mux

// go mod why github.com/gorilla/mux : to why i am dependent on the gorilla mux pkgs 

// go mod graph : it pulls up all the graphs of the dependencies

// go mod edit : to change the mod file
// go mod edit -go 1.17 : to change the version of go

// go mod vendor : give a vendor folder for modules , that comes from the web 