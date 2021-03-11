package main

import (
	"fmt"
	"log"
)

func main() {
	var myString = "Green"
	log.Println("myString is set to", myString)
	changeValue(&myString)
	log.Println("after changing value myString is set to", myString)

}

func changeValue(s *string)  {
	fmt.Println("s is set to ", s)
	newValue := "Red"
	*s = newValue
}
