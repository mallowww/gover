package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Name string `json:"nickname"`
	Age  int
}

func main() {
	data := []byte(`{
		"id": 2,
		"nickname": "okie",
		"age": 15
	}`)

	u := &User{}
	err := json.Unmarshal(data, u)

	fmt.Printf("% #v\n", u)
	fmt.Println(err)
}
