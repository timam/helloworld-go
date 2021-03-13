package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	Date time.Time
}

func main()  {
	user := User{
		FirstName:   "Travor",
		LastName:    "Sawler",
	}

	log.Println(user.FirstName)
}

