package main

import (
	"api"
	"fmt"
)

func main() {
	fmt.Println("hehe")
	hehe := api.ExpStruct{5, 100.0}
	hehe.ModifyIterm(10, 50.0)
	hehe.ShowIterm()
	fmt.Println("nihao")
	a := 10.0
	if a > 20 {
		fmt.Println("haha")
	} else {
		fmt.Println("hehe")
	}

	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}
}
