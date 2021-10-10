package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name string  `json:"coursename"`
	Price int 	
	Platform string `json:"website"`
	Password string  `json:"-"`   //removes the password
	Tags []string   `json:"tags,omitempty"` 
}

func main() {
	fmt.Println("Let's learn about JSON more")
	EncodeJSON()
}

func EncodeJSON(){
	someCourses := []Course{
		{"ReactJS Bootcamp",299,"LCO","abc123",[]string{"web-dev","js"}},
		{"VueJS Bootcamp",599,"AWS","pgq789",[]string{"web-dev","js","vuex"}},
		{"Angular Bootcamp",999,"Coursera","rup435",nil},
	}
	//package this data as json data
	finalJson, err := json.MarshalIndent(someCourses,"","\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}

func DecodeJSON(){
	// uri := "https://api.nuxtjs.dev/mountains"
	// res, err := http.Get(uri)
	// PanicError(err)
	// defer res.Body.Close()
	// content, err := ioutil.ReadAll(res.Body)
	// PanicError(err)
	// stringBuilder := strings.Builder{}
	// length,err := stringBuilder.Write(content)
	// fmt.Println(length)
	// data := stringBuilder.String()
	// fmt.Printf("%T",data)
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LCO",
		"tags": [
				"web-dev",
				"js"
		]
    }
	`)

	var course Course 

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was validated")
		json.Unmarshal(jsonDataFromWeb, &course)
		fmt.Printf("%#v\n",course)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where you want to add data to key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key is %v and value is %v\n", key, value)
	}

}

func PanicError(err error){
if err != nil{
		panic(err)
	}
}