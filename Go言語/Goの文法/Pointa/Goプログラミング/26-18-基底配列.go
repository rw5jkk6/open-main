package main
 
import (
   "fmt"
)

func reclassifiy(planets *[]string){
	*planets = (*planets)[0:3]
}

func main(){
	planets :=[]string{"a", "b", "c", "d", "e"}
	reclassifiy(&planets)
	fmt.Println(planets)  // [a b c]
}