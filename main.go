package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type course struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Department string `json:"Department"`
	CourseNumber string `json:"CourseNumber"`
	Units string `json:"Units"`
	Description string `json:"Description"`
}

type allCourses []course

var courses = allCourses {
	{
		ID: "cs1300",
		Title: "Discrete Structures",
		Department: "Computer Science",
		CourseNumber: "1300",
		Units: "4",
		Description: "Logic, proof techniques, sets, basic counting rules, relations, functions and recursion. Modular arithmetic, prime numbers. Graphs and trees. Boolean algebra, computer logic, and combinational circuits.",
	},
	{
		ID: "cs1400",
		Title: "Introduction to Programming and Problem Solving",
		Department: "Computer Science",
		CourseNumber: "1400",
		Units: "4",
		Description: "Problem-solving methods. Object-oriented concepts including classes, inheritance, and polymorphism. Basic control structures and data types. File I/O and exception handling. Program documentation and testing.",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to your Go home page!")
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	courseID := mux.Vars(r)["id"]

	for _, singleCourse := range courses{
		if singleCourse.ID == courseID {
			json.NewEncoder(w).Encode(singleCourse)
		}
	}
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}