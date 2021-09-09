package main
 
import (
   "fmt"
)

func main(){
	// runeはc言語のcharとは微妙に異なる
	// 32bitは \Uxxxx
	var s rune = '\U0001F628'
	fmt.Println(s) // 128552
	fmt.Printf("%c\n", s)

	var ss string = "\U0001F628"
	fmt.Println(ss)

	fmt.Println("\U0001F600")
}
