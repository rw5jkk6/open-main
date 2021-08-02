package main

import (
	"fmt"
)

type num_data struct{
	a int
	d float32
}

func dealDelta(data num_data){
	fmt.Printf("a=%d f=%f\n", data.a, data.d)
	data.a = 2;
	data.d = 2.4
}

func dealDelta2(data *num_data){
	fmt.Printf("a=%d f=%f\n", data.a, data.d)
	data.a = 2;
	data.d = 2.4
}



func main(){
	n1 := num_data{a: 1, d: 1.2}
	n2 := num_data{a: 1, d: 1.2}

	dealDelta(n1)
	dealDelta2(&n2)

	fmt.Printf("n1.a = %d n1.d = %f\n", n1.a, n1.d)
	fmt.Printf("n2.a = %d n2.d = %f\n", n2.a, n2.d)
}

// a=1 f=1.200000
// a=1 f=1.200000
// n1.a = 1 n1.d = 1.200000
// n2.a = 2 n2.d = 2.400000
