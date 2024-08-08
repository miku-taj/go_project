package main

import "fmt"

type Person struct {
	name string
	age  int
	address Address
}

type Address struct {
	House int
	Street string
	City string
}

type Node struct {
	value int
	next *Node
}

//1
type Age int
//2
type Number int
//3
type Score int
//4
type Operation func(int, int) int
//5
type Comparator func(int, int) bool
//6
type UnaryOperation func(int) int

func main() {

	user1 := Person{name: "Alice", age: 13}
	fmt.Printf("%+v\n", user1)
	fmt.Println("Name is ", user1.name, " Age is ", user1.age)
//	fmt.Println("user1 city", user1.Address.City)

	user2 := new(Person)
	fmt.Printf("%+v", *user2)
	
	user3 := Person{}
	fmt.Println(user3)
	
	user4 := Person{name: "Albus", age: 51}
	fmt.Println(user4.name, " is ", user4.age, " years old")

	n1 := &Node{ value: 1}
	n2 := &Node{ value: 2}
	n3 := &Node{ value: 3}
	n1.next = n2
	n2.next = n3
	current_nodePtr := n1
	for current_nodePtr != nil {
		fmt.Println(*current_nodePtr)
		current_nodePtr = current_nodePtr.next
	}

	//4
	var adder Operation = add
	var subtracter Operation = subtract
	var multiplier Operation = multiply
	fmt.Println("Using adder: ", adder(1, 2))
	fmt.Println("Using subtracter: ", subtracter(3, 2))
	fmt.Println("Using multiplier: ", multiplier(10, 2))

	//5
	var equality_comparator Comparator = is_equal
	var bigger_comparator Comparator = is_larger
	var smaller_comparator Comparator = is_smaller
	fmt.Println("Using equality_comparator ", equality_comparator(10, 10))
	fmt.Println("Using bigger_comparator ", bigger_comparator(1, 10))
	fmt.Println("Using smaller_comparator ", smaller_comparator(2, 10))

	//6
	var twofold UnaryOperation = func (n int) int {
		return n * 2
	}
	var threefold UnaryOperation = func (n int) int {
		return n * 3
	}
	fmt.Println("Using unary operation twofold: ", twofold(3))
	fmt.Println("Using unary operation threefold: ", threefold(3))

	//13
	change_age(&user1, 51)
	//14
	r1 := Rectangle{Width: 10, Height:  15}
	update_area(&r1)
	

}

//1
func is_adult(age Age) string {
	if age >= 18 {
		return "Является совершеннолетним"
	} else {
		return "Является несовершеннолетним"
	}
}

//2
func is_even(n Number) string {
	if n % 2 == 0 {
		return "Является четным"
	} else {
		return "Является нечетным"
	}
}

//3
func is_corect(score Score) string {
	if score >= 0 && score <= 100 {
		return "Находится в допустимом диапазоне"
	} else {
		return "Выходит из допустимого диапазона"
	}
}

//4
func add(a, b int) int{
	return a + b
}
func subtract(a, b int) int{
	return a - b
}
func multiply(a, b int) int{
	return a * b
}

//5
func is_equal(a, b int) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func is_larger(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func is_smaller(a, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
}

//7
type Count int
func countdown(strarting_point Count) {
	for strarting_point >= 0 {
		fmt.Println(strarting_point)
		strarting_point--
	}
}

//8
type Temperature float64
func weather_forecast(temp Temperature) {
	if temp > 0 {
		fmt.Println("Temperature is higher than zero")
	} else if temp == 0 {
		fmt.Println("Temperature is equal to zero")
	} else {
		fmt.Println("Temperature is lower then zero")
	}
}

//9
type Price float64
func sale(starting_price Price) float64 {
	result := starting_price * 0.9
	return float64(result)
}
 
//10
type User struct {
	Name string
	Age int
}
func info_display(user User) {
	fmt.Println("User name: ",user.Name, "\nUser age: ", user.Age)
}

//11
type Product struct{
	Name string
	Price Price
}
func product_price(p Product) {
	fmt.Println("The product cost is ", p.Price)
}

//12
type Book struct {
	Title string
	Author string
	Pages int
}
func about_book(b Book) {
	fmt.Println("The book title is ", b.Title, "\nThe author is ", b.Author, "\nThe number of pages is ", b.Pages)
}

//13
func change_age(p *Person, new_age int) {
	fmt.Println("Changing age of ", p.name)
	fmt.Println("Before: ", *p)
	p.age = new_age
	fmt.Println("After: ", *p)
}

//14
type Rectangle struct{
	Width int
	Height int
	Area int
}

func update_area(r *Rectangle) {
	fmt.Println("Area before: ", r.Area)
	r.Area = r.Height * r.Width
	fmt.Println("Area after: ", r.Area)
}

//15
type Coordinate struct {
	X int
	Y int
}

func compare_coordinates(c1, c2 Coordinate) string {
	if c1 == c2 {
		return "Equal"
	} else {
		return "Not equal"
	}
}