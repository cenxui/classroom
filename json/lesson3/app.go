package main


//WHAT TO DO IF THE FIELD IS EMPTY
//The JSON parser also accepts a flag in the tag to let it know what to do if the field is empty.
// The omitempty flag tells it to not include the JSON value in the output if it’s the “zero-value”
// for that type.
//
//The “zero-value” for numbers is 0, for strings it’s the empty string, for maps, slices and pointers
// it’s nil. This is how you include the omitempty flag.

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
