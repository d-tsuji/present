package main

import (
	"encoding/json"
	"fmt"
)

// json-main OMIT
func main() {
	jsonStr := `{"id":"0001","number":1}`
	var t map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &t); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)
}

// json-main OMIT
