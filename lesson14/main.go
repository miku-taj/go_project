package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	// file, err := os.Create("example.txt")
	// if err != nil{
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()
	// file, err := os.Open("hello.txt")
	// if err != nil{
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()
	// data, err := io.ReadAll(file)
	// if err != nil{
	// 	fmt.Println("Error reading file:", err)
	// }
	// fmt.Println(string(data))
	// _, err = file.Write([]byte("hello"))
	// if err != nil{
	// 	fmt.Println("Error creating file:", err)
	// }

	file, err := os.OpenFile("new_file.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")
	data, error := io.ReadAll(file)
	if error != nil {
		fmt.Println("Error reading file:", error)
	}
	fmt.Println(string(data))
	// _, err = file.Write([]byte("hello"))
	// if err != nil{
	// 	fmt.Println("Error creating file:", err)
	// }

	var input string
	fmt.Fscan(file, &input)
	fmt.Println("input:", input)
}