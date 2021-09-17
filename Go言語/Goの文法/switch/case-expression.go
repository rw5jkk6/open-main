/*
caseに式が書ける。その場合はswitchの中には式は書かない 
*/

package main

import (
	"fmt"
)

func main(){
	day := 5
	// switchと { の間にはtrueがある
	switch {
	case 0 > day && day < 3:
		fmt.Println("0 < day < 3")
	case day > 3 && day < 6:
		fmt.Println("3 < day < 6")
	}
}
