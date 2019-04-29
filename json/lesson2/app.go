package main


// As you might know, Go requires all exported fields to start with a capitalized letter.
// It’s not common to use that style in JSON however. We use the tag to let the parser
// know where to actually look for the value.

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id string `json:"user_id"`
	Title string `json:"user_title"`
}



func main() {

	app := App {
		Id:"12345",
		Title:"good day",
	}
	v, err := json.Marshal(app)

	if err != nil {
		fmt.Errorf(err.Error())
	}
	// value show will be like this {"user_id":"12345","user_title":"good day"}
	// you can see the key value id is to be user_id and title to be user_title
	fmt.Print(string(v))
}
