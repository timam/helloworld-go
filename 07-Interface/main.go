package main

import "log"

type Animal interface {
	 Says() string
	 NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main()  {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shpered",
	}
	printInfo(dog)
}

func (r Dog)Says() string {
	return "woof"
}

func (r Dog)NumberOfLegs() int {
	return 4
}

func printInfo(a Animal)  {
	log.Println("This animal says", a.Says(), "and has ", a.NumberOfLegs(), "legs")
}