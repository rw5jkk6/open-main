package lib

import "testing"

func TestSwap(t *testing.T){
	s := &S{X:1,Y:2}
	s.Swap()
	expect := S{X:2,Y:1}
	if expect != *s{
		t.Errorf("%T, %T", s, expect)
	}
}