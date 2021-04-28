package main

import (
	"fmt"

	"github.com/jcfrancisIII/goMongo/foo"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width += 20
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	a := [2]string{"hello", "world"}
	for i, x := range a {
		fmt.Println("a: ", i)
		fmt.Println("a: ", x)
	}
	fmt.Println(foo.Foo)
	ArrPrint()
}
