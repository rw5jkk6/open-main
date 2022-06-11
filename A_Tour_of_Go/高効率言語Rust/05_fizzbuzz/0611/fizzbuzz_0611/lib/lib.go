package lib

import "strconv"

func FizzBuzz(n int)string{
	if n % 15 ==0{
		return "fizzbuzz"
	}else if n % 3==0{
		return "fizz"
	}else if n % 5 ==0{
		return "buzz"
	}else{
		return strconv.Itoa(n)
	}
}