package main

import (
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 100.0, 20.0, 5.0, false},
}

func TestDivision(t *testing.T)  {
	for _, tt := range tests{
		got, err := devide(tt.dividend, tt.divisor)
		if tt.isErr{
			if err  == nil{
				t.Error("Expected error but did not get one")
			}
		} else {
			if err != nil{
				t.Error("Did not expected error but get one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}

}