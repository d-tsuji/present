package main

import (
	"encoding/json"
	"fmt"
)

// optional-struct OMIT
type user struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Optional string `json:"optional,omitempty"`
}

// optional-struct OMIT

// json-main OMIT
func main() {
	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	b, _ := json.MarshalIndent(u, "", "   ")
	fmt.Println(string(b))
}

// json-main OMIT
