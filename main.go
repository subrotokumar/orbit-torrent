package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	value := 134
	json, _ := json.Marshal(value)
	fmt.Println(json)
}
