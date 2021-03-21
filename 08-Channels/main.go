package main

import (
	"github.com/timam/helloworld-go/08-Channels/helpers"
	"log"
)

const numPool = 10

func calculateValue(intChan chan int)  {
	randomNumber := helpers.RandormNumber(numPool)
	intChan <- randomNumber
}

func main()  {
	intChan := make(chan int)
	defer close(intChan)
	go calculateValue(intChan)
	num := <- intChan
	log.Println(num)
}