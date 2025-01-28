package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"authorName"`
	Website    string `json:"website"`
}

func (c Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

var courses []Course

func main() {
	fmt.Println("Build API with Sajib")
	r := mux.NewRouter()

	courses = append(courses,
		Course{
			CourseId:    "1",
			CourseName:  "ReactJS",
			CoursePrice: 299,
			Author: &Author{
				AuthorName: "Hitesh Choudhary",
				Website:    "lco.dev",
			},
		},
		Course{
			CourseId:    "2",
			CourseName:  "MERN Stack",
			CoursePrice: 199,
			Author: &Author{
				AuthorName: "Hitesh Choudhary",
				Website:    "go.dev",
			},
		},
	)

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", getSingleCourse).Methods("GET")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PATCH")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/course/partialUpdate/{id}", updatePartialCourse).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by SajibUzzaman</h1>"))
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		// json.NewEncoder(w).Encode("Please send some data")
		http.Error(w, "Please send some data", http.StatusBadRequest)
		return
	}

	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid Json formate", http.StatusBadRequest)
		return
	}

	if course.IsEmpty() {
		// json.NewEncoder(w).Encode("No data inside JSON")
		http.Error(w, "No data inside JSON", http.StatusBadRequest)
		return
	}

	// Duplicate title check
	for _, item := range courses {
		if item.CourseName == course.CourseName {
			// json.NewEncoder(w).Encode("This course name is already exists please try another")
			http.Error(w, "This course name already exists, please try another", http.StatusConflict)
			return
		}
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")

	if courses == nil {
		courses = []Course{}
	}

	json.NewEncoder(w).Encode(courses)
}

func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Single Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]

	for _, course := range courses {
		if course.CourseId == courseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// json.NewEncoder(w).Encode("No course found with this given id")
	http.Error(w, "No course found with this given id", http.StatusNotFound)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]

	for index, item := range courses {
		if item.CourseId == courseId {
			// courses = append(courses[:index], courses[index+1:]...)

			// var updatedCourse Course
			// if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
			// 	http.Error(w, "Invalid Json formate", http.StatusBadRequest)
			// 	return
			// }
			// updatedCourse.CourseId = courseId
			// courses = append(courses, updatedCourse)

			var updatedCourse Course
			if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
				http.Error(w, "Invalid Json formate", http.StatusBadRequest)
				return
			}

			updatedCourse.CourseId = courseId
			courses[index] = updatedCourse

			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	// json.NewEncoder(w).Encode("No course found with this given id")
	http.Error(w, "No course found with this given id", http.StatusNotFound)
}

func updatePartialCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Partial update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]

	for index, item := range courses {
		if item.CourseId == courseId {

			var partialUpdate map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&partialUpdate); err != nil {
				http.Error(w, "Invalid Json formate", http.StatusBadRequest)
				return
			}

			if courseName, exists := partialUpdate["courseName"]; exists {
				item.CourseName = courseName.(string)
			}

			if price, exists := partialUpdate["price"]; exists {
				item.CoursePrice = int(price.(float64))
			}

			if authorData, exists := partialUpdate["author"]; exists {
				authorMap := authorData.(map[string]interface{})

				if authorName, exists := authorMap["authorName"]; exists {
					item.Author.AuthorName = authorName.(string)
				}
				if website, exists := authorMap["website"]; exists {
					item.Author.Website = website.(string)
				}
			}

			courses[index] = item

			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// json.NewEncoder(w).Encode("No course found with this given id")
	http.Error(w, "No course found with this given id", http.StatusNotFound)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]

	for index, item := range courses {
		if item.CourseId == courseId {
			courses = append(courses[:index], courses[index+1:]...)

			json.NewEncoder(w).Encode("Successfully Delete")
			return
		}
	}

	// json.NewEncoder(w).Encode("No course found with this given id")
	http.Error(w, "No course found with this given id", http.StatusNotFound)
}
