package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct)printFirstName() string {
	return m.FirstName
}

func main()  {
	var myvar  myStruct
	myvar.FirstName = "Jhon"

	myvar2 := myStruct{
		FirstName: "Mary",
	}
	log.Println(myvar.printFirstName())
	log.Println(myvar2.printFirstName())
}