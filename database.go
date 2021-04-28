package main

import "fmt"

func ArrPrint() {
	a := [2]uint64{1, 1}
	for i, x := range a {
		fmt.Println("key: ", i)
		fmt.Println("val: ", x)
	}
}
