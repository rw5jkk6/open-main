package main

import (
	"testing"
)

func TestFizzbuzz(t *testing.T){
	tests := []struct{
		expect string
		actual int
	}{
		{"fizz", 3},
		{"buzz", 5},
		{"fizzbuzz", 15},
		{"buzz", 50},
	}

	for _ , v := range tests{
		if v.expect != Fizzbuzz(v.actual){
			t.Errorf("%s != %d", v.expect, v.actual)
		}
	}
}


// func TestFizzbuzz(t *testing.T){
// 	expect := "fizzbuzz"
// 	actual := Fizzbuzz(15)
	
// 	if expect != actual{
// 		t.Errorf("%s != %s", expect, actual)
// 	}
// }