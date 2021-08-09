package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//Поля структуры должны быть экспортируемые если
	//мы хотим чтобы пакет JSON с ними мог работать
	type CrewwMember struct {
		ID                int      `json:"id,omitempty"`
		Name              string   `json:"name"`
		SecurityClearance int      `json:"clearanceLevel"`
		AccessCodes       []string `json:"accesscodes"`
	}

	type StarShip struct {
		ShipID    int
		ShipClass string
		Capitan   CrewwMember
	}

	sbyte := []byte(`[{"id":1,"name":"Jacob","clearanceLevel":10,"accesscodes":["ADA","LAL"]},{"id":2,"name":"Coby","clearanceLevel":15,"accesscodes":["TTL","ALA"]}]`)
	s := []CrewwMember{}

	json.Unmarshal(sbyte, &s)

	fmt.Println(s)
}
