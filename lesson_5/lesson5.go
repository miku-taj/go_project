package main

import "fmt"

func main() {
	var x int = 10
	var p *int
	p = &x

	fmt.Println("x: ", x)
	fmt.Println("x address: ", &x)
	fmt.Println("p: ", p)
	fmt.Println("p points to: ", *p)

	//1
	fmt.Println("Before: ", x)
	updateValue(p)
	fmt.Println("After: ", x)
}

//1
func updateValue(p *int) {
	*p =  *p + 10
}

//2
