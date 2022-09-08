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

type Student struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Grade     int32  `json:"grade"`
	Major     string `json:"major"`
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	// Set the Header Content-type as application/json
	w.Header().Set("Content-type", "application/json")

	//Declare a student variable
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(1000000000))
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
		}
	}
}

func main() {
	students = append(students, Student{ID: "1", Firstname: "Max", Lastname: "Tizie", Grade: 90, Major: "Computer Science"})
	r := mux.NewRouter()
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")

	fmt.Printf("Starting server on port 8282\n")
	log.Fatal(http.ListenAndServe(":8282", r))

}
