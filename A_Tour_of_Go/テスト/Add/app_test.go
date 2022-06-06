package add

import (
	"testing"
)

func TestAdd(t *testing.T){
	expect := 3
	actual := Add(1, 2)
	if actual != expect{
		t.Error("%S != %S", actual, expect)
	}
}