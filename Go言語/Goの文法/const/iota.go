package main
 
import (
   "fmt"
)

type Month int

const (
	January   Month = 1 + iota
	February 
	March
)

const (
	_ = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main(){
	fmt.Println(March)
	fmt.Println(KB, MB, GB)
}

// 3
// 1024 1048576 1073741824
