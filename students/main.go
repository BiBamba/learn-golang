package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func main() {
	students = append(students, Student{ID: "1", Firstname: "Max", Lastname: "Tizie", Grade: 90, Major: "Computer Science"})
	r := mux.NewRouter()
	r.HandleFunc("/students", getStudents).Methods("GET")
	//r.HandleFunc("/student/{id}", getStudent).Methods("GET")

	fmt.Printf("Starting server on port 8282\n")
	log.Fatal(http.ListenAndServe(":8282", r))

}
