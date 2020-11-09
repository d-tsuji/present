package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// json-struct OMIT
// JSONに共通的に含めたい項目を構造体として埋め込む
type Record struct {
	ID string `json:"id"`
	Common
}

type Common struct {
	CreateAt time.Time `json:"create_at"`
}

// json-struct OMIT

// json-main OMIT
func main() {
	r := Record{
		ID:     "hoge",
		Common: Common{CreateAt: time.Now()},
	}

	b, _ := json.MarshalIndent(r, "", "   ")
	fmt.Printf("%+v", string(b))
}

// json-main OMIT
