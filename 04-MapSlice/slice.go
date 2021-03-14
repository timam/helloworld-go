package main

import "log"

func main()  {
	mySlice := []string{}
	mySlice = append(mySlice,"Timam")
	mySlice = append(mySlice,"Talukdar")
	log.Println(mySlice)
	numbers := []int{1,2,3,4,5,6,7,8,9}
	log.Println(numbers[4:9])
}
