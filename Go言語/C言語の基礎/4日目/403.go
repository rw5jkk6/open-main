package main
 
import (
   "fmt"
)

func avg(m float64, n float64) float64{
	return (m + n) / 2
}

func main(){
	var d1, d2, d3 float64
	
	/* 
	この場合、型がつけられないので、
	デフォルトがfloat64なので、それで関数を作る
	*/
	var (
		a = 1.2
		b = 3.4
		c = 2.7
	)

	d1 = avg(a, b)
	d2 = avg(4.1, 5.7)
	d3 = avg(c, 2.8)
	fmt.Printf("d1=%f d2=%f d3=%f \n", d1, d2, d3)
  // d1=2.300000 d2=4.900000 d3=2.750000 
}
