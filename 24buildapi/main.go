package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Model fir course - file

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fakeDB
var courses []Course

//middleware - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}


func main() {
	fmt.Println("Welcome to courses")
}

//controllers - file
//serve home route

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API By Rupam Shil</h1>"))
} 

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All course")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}