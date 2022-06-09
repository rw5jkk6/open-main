package main

import (
	"testing"
)

func TestSwap(t *testing.T){
	actual := 2
	s := SwapStr{a:1, b:2}
	s.Swap()
	expect := s.a
	if actual != expect{
		t.Error("%S != %S", actual, expect)
	}
}