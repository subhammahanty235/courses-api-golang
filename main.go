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

// API development

type Course struct {
	CourseId    string  `json:"courseid`
	CourseName  string  `json:"coursename`
	CoursePrice int     `json:price`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// DB

var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// controllers

func homecontroller(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from <h1>Server</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send Some Data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data provided")
		return
	}

	// generate unique id
	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}

	}

}

func main() {
	r := mux.NewRouter()

	// seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "Frontend Development", CoursePrice: 299, Author: &Author{Fullname: "Subham", Website: "https://Com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Backend Development", CoursePrice: 599, Author: &Author{Fullname: "Subham", Website: "https://Com"}})
	courses = append(courses, Course{CourseId: "5", CourseName: "Full Stack Development", CoursePrice: 899, Author: &Author{Fullname: "Subham", Website: "https://Com"}})

	r.HandleFunc("/", homecontroller).Methods("GET")
	r.HandleFunc("/getcourses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/updatecourse/{id}", updateOneCourse).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5000", r))

}
