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

	sbyte := []byte(`{"ShipID":3,"ShipClass":"Fighter","Capitan":{"id":1,"name":"Jacob","clearanceLevel":10,"accesscodes":["ADA","LAL"]}}`)
	si := new(StarShip)

	json.Unmarshal(sbyte, si)

	fmt.Println(string(si.ShipClass))
}
