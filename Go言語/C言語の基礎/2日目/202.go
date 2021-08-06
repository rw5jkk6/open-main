package main

import (
	"fmt"
)

func main(){
	fmt.Printf("私の名前%sはです。年齢は%d歳です。\n", "山田太郎", 20)
	fmt.Printf("イニシャルは%cです。\n", 'Y')
	// fmt.Printf("イニシャルは%sです。\n", "Y") // 'Y'だとアスキーコードになる
	// fmt.Printf("イニシャルは%vです。\n", "Y")
	fmt.Printf("%f + %f = %f\n", 1.2, 2.7, 1.2 + 2.7)
}

// 私の名前山田太郎はです。年齢は20歳です。
// イニシャルはYです。
// 1.200000 + 2.700000 = 3.900000
