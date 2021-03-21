package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name" : "Jhon",
		"last_name" : "Wick",
		"hair_color" : "black",
		"has_dog" : true
	},
	{
		"first_name" : "Bruce",
		"last_name" : "Wayne",
		"hair_color" : "black",
		"has_dog" : false
	}
]
`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	fmt.Printf("unmarshalled: %v\n", unmarshalled)

	//write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Raju"
	m1.LastName = "Sarkar"
	m1.HairColor = "Red"
	m1.HasDog = false
	mySlice = append(mySlice, m1)

	var m2 Person
	m1.FirstName = "Esha"
	m1.LastName = "Sarkar"
	m1.HairColor = "Blue"
	m1.HasDog = true
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice,"","    ")
	if err != nil{
		log.Println("Error Marshaling newJson", err)
	}
	fmt.Println(string(newJson))
}
