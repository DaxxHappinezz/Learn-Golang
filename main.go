package main

import (
	"fmt"

	"github.com/DaxxHappiness/learnGo/mydict"
)


func main() {
	dictionary := mydict.Dictionary{}
	hello := "Hello"
	err := dictionary.Add(hello, "First")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} 
	val, _ := dictionary.Search(hello)
	fmt.Printf("val: %v\n", val)
	err = dictionary.Update("lsekjflksejf", "second")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	val, _ = dictionary.Search(hello)
	fmt.Printf("val: %v\n", val)
}