package main

import (
	"encoding/json"
	"fmt"
)

// json-main OMIT
func main() {
	u := User{
		UserID: "0001",
		Name:   "gopher",
	}
	b, _ := json.MarshalIndent(u, "", "   ")
	fmt.Printf("%+v\n", string(b))
}

// json-main OMIT

// json-struct OMIT
type User struct {
	UserID    string   `json:"user_id"`
	Name      string   `json:"user_name"`
	Languages []string `json:"languages"`
}

// json-struct OMIT
