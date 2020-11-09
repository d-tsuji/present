package main

import (
	"encoding/json"
	"fmt"
)

// json-struct OMIT
type TypeRequest struct {
	Type   string          `json:"type"`
	Result json.RawMessage `json:"result"`
}

type ID struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// json-struct OMIT

// json-input OMIT
var typeIDJSON = `{"type":"1","result":{"id":"1","name":"name"}}`

// json-input OMIT

// json-main OMIT
func main() {
	var r TypeRequest
	if err := json.Unmarshal([]byte(typeIDJSON), &r); err != nil {
		panic(err)
	}

	// デコード済のフィールドを用いて残りの構造を決める
	switch r.Type {
	case "1":
		var id ID
		if err := json.Unmarshal(r.Result, &id); err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", id)
	case "2":
		// ...
	default:
		panic("cannot handle json")
	}
}

// json-main OMIT
