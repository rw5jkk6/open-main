package main

import "fmt"

// 素数ならtrueを返す
func Is_Prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Get_Primes(primes []int) {
	count := 0
	for i := 2; count < 100; i++ {
		if Is_Prime(i) {
			fmt.Println(i)
			primes[count] = i
			count += 1
		}
	}
}
