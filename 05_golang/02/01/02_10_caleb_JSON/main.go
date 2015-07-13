package main
import (
	"encoding/json"
	"fmt"
	"bytes"
)

func main() {
	jsonData := `
	{
	100
	}
	`

	var obj interface{}

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	bs, err := json.NewEncoder(&buf).Encode([]int{1,2,3,4})
	fmt.Println(string(bs), err)
}
