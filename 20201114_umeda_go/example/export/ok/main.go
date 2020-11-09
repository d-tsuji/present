package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(err)
	}
	var r ip
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", r)
}

// json-struct-ip OMIT
type ip struct {
	Origin string `json:"origin"`
}

// json-struct-ip OMIT
