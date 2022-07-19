package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// create model for courses -file

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
// using slice
var courses []Course

// middleware or helper - file

func (c *Course) IsEmpty() bool {

	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("Welcome to build the API in golang")
	r := mux.NewRouter()
	// seeding of data

	courses = append(courses, Course{CourseID: "2", CourseName: "golang", CoursePrice: 299, Author: &Author{Fullname: "sumit chahar", Website: "www.sc.com"}})

	courses = append(courses, Course{CourseID: "4", CourseName: "cpp", CoursePrice: 199, Author: &Author{Fullname: "sumit Choudhary", Website: "www.schoudhary.com"}})

	// routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/coursecreate", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to the port
	log.Fatal(http.ListenAndServe(":8000", r))

}

// controllers - file
// server home route

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>WELCOME TO API ON GOLANG </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from reques
	params := mux.Vars(r)

	// fmt.Println(params)

	// loop through courses , find the matching id and return the response

	for _, course := range courses {

		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("no course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// entire body is empty then
	if r.Body == nil {

		json.NewEncoder(w).Encode("please send some data")
	}

	// data send up like {}
	var cour Course
	_ = json.NewDecoder(r.Body).Decode(&cour)

	if cour.IsEmpty() {

		json.NewEncoder(w).Encode("no data inside the json")
		return

	}

	// to generate unique id, and convert it to string
	// append new course into courses

	rand.Seed(time.Now().UnixNano())
	cour.CourseID = strconv.Itoa(rand.Intn(100))

	courses = append(courses, cour)
	json.NewEncoder(w).Encode(cour)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from request
	params := mux.Vars(r)

	// loop through the value , once get the ID and remove it from the index and add value back in the courses with myID

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var cours Course
			_ = json.NewDecoder(r.Body).Decode(&cours)

			cours.CourseID = params["id"]
			courses = append(courses, cours)

			json.NewEncoder(w).Encode(cours)
			return
		}
	}
	//TODO: send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop , hit the id, remove operation

	for index, coursee := range courses {
		if coursee.CourseID == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)

			break

			// json.NewDecoder(r.Body).Decode(&coursee)

		}
	}
}
