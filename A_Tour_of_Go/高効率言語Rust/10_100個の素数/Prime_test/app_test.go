package main

import (
	"testing"
)

func TestIs_Prime(t *testing.T){
	if true != Is_Prime(7){
		t.Errorf("%v != %d", true, 7)
	}
}