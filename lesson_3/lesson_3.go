package main

import "fmt"

func main() {

	num := 6

	if num > 9 && num < 100 {
		fmt.Println("Двузначное")
	}

	if num % 2 == 1 {
		fmt.Println("Число нечетное")
	} else {
		fmt.Println("Число четное")
	}

	number := 100

	if number >=0 && number < 10 {
		fmt.Println("Однозначное")
	} else if number > 9 && number < 100 {
		fmt.Println("Двузначное")
	} else if number > 99 && number < 1000 {
		fmt.Println("Трехзначное")
	} else if number > 999 && number < 10000 {
		fmt.Println("Четырехзначное")
	} else {
		fmt.Println("Число слишком большое")
	}

	var a, b int = 10, 22
	if a >= b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	c, d, e := 3, 4, 5

	if c >= d {
		if d >= e {
			fmt.Println(c, d, e)
		} else {
			fmt.Println(c, e, d)
		}
	} else {
		if c >= e {
			fmt.Println(d, c, e)
		} else {
			fmt.Println(d, e, c)
		}
	}

	nn := 1

	switch nn {
	case 1 :
		fmt.Println("Number is positive")
		fallthrough
		//fallthrough makes the next case block to be executed although the condition is false 
	case -1 :
		fmt.Println("Number is negative")
	default :
		fmt.Println("Number is unsigned")
	}

	n1, n2, n3 := 10, 33,2
	sum := 0
	if n1 >= n3 && n2 >= n3 {
		sum = n1 + n2
	} else if n1 >= n2 && n3 >= n3 {
		sum = n1 + n3
	} else if n3 >= n1 && n2 >= n1 {
		sum = n3 + n2
	}
	fmt.Println(sum)

	year := 2000

	if year % 4 == 0 {
		if year % 100 == 0 && year % 400 != 0 {
			fmt.Println(365)
		} else {
			fmt.Println(366)
		}
	} else {
		fmt.Println(365)
	}

	month := 2
	switch(month) {
	case 1: fmt.Println(31)
	case 2: fmt.Println(28)
	case 3: fmt.Println(31)
	case 4: fmt.Println(30)
	case 5: fmt.Println(31)
	case 6: fmt.Println(30)
	case 7: fmt.Println(31)
	case 8: fmt.Println(31)
	case 9: fmt.Println(30)
	case 10: fmt.Println(31)
	case 11: fmt.Println(30)
	case 12: fmt.Println(31)

	}


}