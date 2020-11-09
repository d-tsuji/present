package main

import (
	"encoding/json"
	"fmt"
)

// json-main OMIT
func main() {
	jsonStr := `{"id":"0001","number":1}`
	var r Record
	if err := json.Unmarshal([]byte(jsonStr), &r); err != nil {
		panic(err)
	}

	fmt.Printf("record=%+v\n", r)
}

// json-main OMIT

// json-struct OMIT
type Record struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
}

// json-struct OMIT
