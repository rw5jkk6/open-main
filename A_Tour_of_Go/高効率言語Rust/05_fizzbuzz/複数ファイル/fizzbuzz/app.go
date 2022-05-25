package main

import "strconv"

func Fizzbuzz(n int)string{
	if n % 3 == 0{
		return "fizz"
	}else if n % 5 == 0{
		return "buzz"
	}else if n % 15 == 0{
		return "fizzbuzz"
	}else{
		return strconv.Itoa(n)
	}
}
