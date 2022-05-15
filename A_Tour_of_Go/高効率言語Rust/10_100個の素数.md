```go
package main
 
import (
   "fmt"
)

func is_Prime(n int)bool{
	for i := 2; i < n; i++{
		if n % i == 0{
			return false
		}
	}
	return true
}

func get_Primes(primes []int){
	count := 0
	for i:=2; count < 100; i++{
		if is_Prime(i){
			fmt.Println(i)
			primes[count] = i
			count += 1
		}
	}
}

func main(){
	primes := make([]int,100)
	get_Primes(primes)
	fmt.Println(primes)
}
```
