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

	str, ok := t["id"].(string)
	if !ok {
		panic("cannot convert")
	}
	num, ok := t["number"].(float64)
	if !ok {
		panic("cannot convert")
	}

	fmt.Printf("str=%v, num=%v\n", str, num)
}

// json-main OMIT
