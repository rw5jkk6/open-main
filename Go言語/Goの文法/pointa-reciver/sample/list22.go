// softwera Desigh 特集1章　cf 25

package main
 
import (
   "fmt"
)

type User struct{
	Name string
	Age int
}

func (u *User)Aging(){
	u.Age++
}

func (u User)AgingButNot(){
	u.Age++
}

func main(){

	u := &User{
		Name: "Richard",
		Age: 33,
	}
	
	u.Aging()
	fmt.Println(u.Age)   // 34

	u.AgingButNot()
	fmt.Println(u.Age)   // 34
}