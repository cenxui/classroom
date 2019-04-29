package main

// The field with the name "-" is also no Unmarshal value
//

import (
	"encoding/json"
	"fmt"
)


type App struct {
	Id string `json:"-"`
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

	//print struct show that { My Awesome App}
	fmt.Println(app)
}
