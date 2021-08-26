package main

import (
	"fmt"
)

func serina()int{
	return 20
}

func hitoshi(len int, s func()int){
	l := s()
	fmt.Printf("hitoshiは%dでserinaは%dです\n", len, l)
}

func main(){
	hitoshi(3, serina)
}

// hitoshiは3でserinaは20です
