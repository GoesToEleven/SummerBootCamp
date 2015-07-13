package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Anything interface{}

func main() {
	jsonData := `
"100"  `
	var obj interface{}
	var f io.Reader
	json.NewDecoder(f).Decode(&obj)

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	// v, ok := obj.(float64)
	// if !ok {
	// 	v = 0
	// }
	// x := 100 + v

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode([]int{1, 2, 3, 4})

	//bs, err := json.Marshal()
	fmt.Println(string(bs), err)

}



