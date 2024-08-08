package main

import "fmt"

func main() {
	//n := create_node()
	a := []int{0, 1, 2}
	s := []int{0, 1}
	
	b := "abcd"
	fmt.Println(b[:2])

	c := make([]int, 5)
	d := make([]int, 5, 10)
	fmt.Println(c, "\n", d)
	fmt.Println(len(append(s, a...)), cap(append(s, a...)))
	
//copy inserts src elements as long as there is place - if cap(dst) is 2, then only 2 el from 
//src will be inserted. After insertion follow rest of the elements dst had before
	src := []int{0, 3, 4}
	dst := []int{1, 2, 10, 11, 12}
	n:= copy(dst, src)
	fmt.Println(dst,"\n", n)


	//1
	arr_int := [7]int{1, 2, 3, 4, 5, 6, 7}
	//2
	arr_str := [5]string{"a", "b", "c", "d", "e"}
	//3
	arr_bool := [4]bool{true, false, true, false}
	fmt.Println(arr_int, arr_str, arr_bool)
	//4
	arr := [10]int{}
	fmt.Println(arr)
	//5
	fmt.Println(arr_bool[1], arr_bool[3])
	//6
	arr2 := [3]bool{true, false, false}
	arr2[0] = false
	//7
	arr3 := [4]string{"ey", "gee", "no", "pe"}
	for i := range arr3 {
		fmt.Println(i)
	}
	//array 19
	array := [10]int{2, 3, 4, 5, 6, 7, 8, 5, 6}
	id := 0
	for i := 1; i < len(array) - 1; i++ {
		if array[i] > array[0] && array[i] < array[9] {
			fmt.Println(id)
			id = i
		}
	}
	fmt.Println(id)

	//
}

type Node struct {
	Value int
	Next  *Node
}

func create_node(value ...int) Node {
	p := new(Node)
	if len(value) > 0 {
		p.Value = value[0]
	}
	return *p
}

func (n Node) end() Node {
	for n.Next != nil {
		n = *n.Next
	}
	return n
}
