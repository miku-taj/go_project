package main

import "fmt"

var Global_var int = 1

func main() {
	var a int8
	var b int16
	var c int32
	var d int64
	println(a, b, c, d)

	var real_number1 float32
	var real_number2 float32
	fmt.Println(real_number1, real_number2)

	var is_true bool
	fmt.Printf("%t is a boolean value", is_true)
	is_true = true

	var a_byte byte
	var a_rune rune

	fmt.Println(a_byte, a_rune)
	var str string = "Hello Go"
	print(str)
	//2

	var num1, num2 int = 2, 4
	fmt.Printf("%d + %d = 6", num1, num2)

}