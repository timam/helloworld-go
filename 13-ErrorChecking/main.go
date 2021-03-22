package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintf(w, "this is the home page")
}
func About(w http.ResponseWriter, r *http.Request)  {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

func addValues(x, y int) int {
	return x+y
}

func divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0,0.0)
	if err != nil{
		fmt.Fprintf(w, "Cant divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divied by %f is %f", 100.0,0.0, f))
}
func divideValues(x,y float32) (float32, error)  {
	if y <= 0{
		err:= errors.New("Cant divide by 0")
		return 0, err
	}
	result := x/y
	return result, nil
}

func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", divide)

	fmt.Println(fmt.Sprintf("Application started at port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
