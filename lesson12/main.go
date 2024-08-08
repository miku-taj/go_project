package main

import "fmt"

func main() {
	// 4 вида инициализации 
	//1
	var DefaultMap map[int]int
	n, exists := DefaultMap[0]
	if exists {
		fmt.Println(n)
	} else {
		fmt.Println("Element not found")
	}
	//2
	MyMap1 := make(map[int]int, 0)
	MyMap2 := make(map[int]int, 10)
	fmt.Println("MyMap1 length: ", MyMap1, "MyMap2 length: ", MyMap2)

	//3
	NamesAgeMap := map[string]int{"Dasha": 5, "Lora": 20}
	fmt.Printf("Type: %T, Value: %#v, Length: %d", NamesAgeMap, NamesAgeMap, len(NamesAgeMap))

	//4
	NewMap := *new(map[string]string)
	fmt.Println(NewMap)

	//удалить эл карты по ключу
	fmt.Println("Before: ", NamesAgeMap)
	delete(NamesAgeMap, "Lora")
	fmt.Println("After: ", NamesAgeMap)

	//
	original := map[string]int{"key1": 10, "key2": 20}
	//copy := original
	copy := make(map[string]int)
	for k, v := range original {
		copy[k] = v
	}
	delete(original, "key1")
	fmt.Println(original, copy) // map[key1:10 key2:20]
	


}
