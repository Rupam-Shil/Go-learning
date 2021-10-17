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
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" 
}


func main() {
	fmt.Println("API Learning")
	r := mux.NewRouter()
	//seeding
	courses = append(courses,Course{CourseId:"2",CourseName:"React Js",CoursePrice:299,Author:&Author{Fullname:"Rupam SHil",Website:"https://github.com/"}})
	courses = append(courses,Course{CourseId:"4",CourseName:"Anuglar Js",CoursePrice:199,Author:&Author{Fullname:"SUbhodeep dey ",Website:"https://twitter.com/"}})
	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCOurse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

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

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/json")
	//grab id from req
	params := mux.Vars(r)

	for _,course := range courses{
		if params["id"] == course.CourseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Welcome to inserting course in database")
	w.Header().Set("Content-Type","application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//check if title is duplicate
	for _, title := range courses {

		if course.CourseName == title.CourseName {
			json.NewEncoder(w).Encode("Duplicate course found")
			return
		}
	}

	// generate a unique id, convert into string

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from req
	params := mux.Vars(r)
	//loop, id and remove the value and add the value again in the courses with myID
	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("could not find course")
}

func deleteOneCOurse( w http.ResponseWriter, r *http.Request){
	fmt.Println("Time to delete some courses")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// if r.Body == nil {
	// 	json.NewEncoder(w).Encode("Please send some data")
	// }

	// var course Course
	// _ = json.NewDecoder(r.Body).Decode(&course)
	// if course.IsEmpty() {
	// 	json.NewEncoder(w).Encode("THe json data is empty")
	// }

	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully deleted an item")
			break
		}
	}
	json.NewEncoder(w).Encode("COuldn't found the data")
}