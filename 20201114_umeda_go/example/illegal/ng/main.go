package main

import (
	"encoding/json"
	"fmt"
)

// optional-struct OMIT
type user struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Fn       func()
}

// optional-struct OMIT

// json-main OMIT
func main() {
	u := user{
		UserID:   "001",
		UserName: "gopher",
		Fn:       func() { fmt.Println("dummy") },
	}
	b, err := json.MarshalIndent(u, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

// json-main OMIT
