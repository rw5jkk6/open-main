package main

import (
	"testing"
)

func TestFizzbuzz(t *testing.T){
	// test1
	expect := "fizzbuzz"
	actual := Fizzbuzz(15)
	
	if expect != actual{
		t.Errorf("%s != %s", expect, actual)
	}

	// test2
	expect = "fizz"
	actual = Fizzbuzz(3)
	
	if expect != actual{
		t.Errorf("%s != %s", expect, actual)
	}

	// test3
	expect = "7"
	actual = Fizzbuzz(7)
	
	if expect != actual{
		t.Errorf("%s != %s", expect, actual)
	}
}	
