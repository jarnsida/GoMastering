package main

import (
	"fmt"

	"github.com/jarnsida/GoMastering/classFactoryTutorial/applience"
)

func main() {
	var myType int

	fmt.Println("Enter preffered applience type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	fmt.Scan(&myType)

	myApplience, err := applience.CreateApplience(myType)

	if err == nil {
		myApplience.Start()
		fmt.Println(myApplience.GetPurpose())
	} else {
		fmt.Println(err)
	}

}
