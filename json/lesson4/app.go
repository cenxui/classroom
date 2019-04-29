package main


//SKIPPING FIELDS
//To have the JSON parser/writer skip a field, just give it the name "-".

import (
	"encoding/json"
	"fmt"
)

type App1 struct {
	Id string `json:"user_id"`
	Title string `json:"user_title"`
}

// Notice that the flag goes inside the quotes.
type App2 struct {
	Id string `json:"user_id,omitempty"`
	Title string `json:"user_title,omitempty"`
}



func main() {

	app1 := App1 {
	}
	v, err := json.Marshal(app1)

	if err != nil {
		fmt.Errorf(err.Error())
	}
	// value show will be like this {"user_id":"","user_title":""}
	fmt.Println(string(v))

	app2 := App2 {
	}

	v, err = json.Marshal(app2)

	if err != nil {
		fmt.Errorf(err.Error())
	}
	// value show will be like this {}
	fmt.Println(string(v))


}
