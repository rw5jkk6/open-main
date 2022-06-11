package lib

import (
	"testing"
)

func Testfizzbuzz(t *testing.T){
	fb := []struct{
		act int
		exp string
	}{
		{3, "fizz"},
		{5, "buzz"},
		{15, "fizzbuzz"},
		{4, "4"},
	}
	for _ , tt := range fb{
		if tt.exp != FizzBuzz(tt.act){
			t.Errorf("%s != %d", tt.exp, tt.act)
		}
	}
}