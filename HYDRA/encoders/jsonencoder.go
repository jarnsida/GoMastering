package main

import (
	"encoding/json"
	"log"
	"os"
)

type CrewwMember struct {
	ID                int      `json:"id,omitempty"`
	Name              string   `json:"name"`
	SecurityClearance int      `json:"clearanceLevel"`
	AccessCodes       []string `json:"accesscodes"`
}

type ShipInfo struct {
	ShipID    int
	ShipClass string
	Capitan   CrewwMember
}

func main() {
	f, err := os.Create("jfile.json")
	PrintFatalError(err)

	defer f.Close()

	//create some data to encode
	cm := CrewwMember{1, "Jacob", 10, []string{"ADA", "LAL"}}
	si := ShipInfo{1, "Fighter", cm}

	err = json.NewEncoder(f).Encode(&si)
	PrintFatalError(err)

}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened during processing file", err)
	}
}
