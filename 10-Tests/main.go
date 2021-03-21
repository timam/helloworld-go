package main

import (
	"errors"
	"log"
)

func main()  {
	result, err := devide( 100.0, 10)
	if err != nil{
		log.Println(err)
		return
	}
	log.Println("divide result is ", result)
}

func devide(x, y float32) (float32, error) {
	var result float32

	if y == 0{
		return result, errors.New("can not divide by zero")
	}

	result = x / y
	return result, nil
}