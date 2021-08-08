package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type CrewwMember struct {
		ID                int32
		Name              string
		SecurityClearance int32
	}

	cm := CrewwMember{1, "Jacob", 10}

	b, err := json.Marshal(cm)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
}
