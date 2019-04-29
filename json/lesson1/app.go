package main

//Parsing a format like JSON in a statically typed language like Go presents a bit of a problem.
// If anything could show up in the JSON body, how does the compiler know how to setup memory to
// have a spot to place everything?
//
//There are two answers to this. The easy option, for when you know what your data will look like,
// is to parse the JSON into a struct you’ve defined. Any field which doesn’t fit in the struct will
// just be ignored. We’ll cover that option first.

import (
	"encoding/json"
	"fmt"
)


type App struct {
	Id string `json:"id"`
	Title string `json:"title"`
}


func main() {

	data := []byte(`
    {
        "id": "k34rAT4",
        "title": "My Awesome App"
    }
`)

	var app App
	//make json to struct
	err := json.Unmarshal(data, &app)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	//print struct show that {k34rAT4 My Awesome App}
	fmt.Println(app)

	v, err := json.Marshal(app)

	if err != nil {
		fmt.Errorf(err.Error())
	}
	//print json show that {"id":"k34rAT4","title":"My Awesome App"}
	fmt.Println(string(v))
}
