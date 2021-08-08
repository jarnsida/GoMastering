package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type CrewwMember struct {
		ID                int      `json:"id,omitempty"`
		Name              string   `json:"name"`
		SecurityClearance int      `json:"clearanceLevel"`
		AccessCodes       []string `json:"accesscodes"`
	}

	cm := CrewwMember{1, "Jacob", 10, []string{"ADA", "LAL"}}

	b, err := json.Marshal(&cm)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
}
