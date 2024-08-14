package main

import "fmt"

type UserInfo interface{
	GetUserInfo() string
	GetUserName() string
	GetUserAge() int
}

type User struct {
	Name string
	Age  int
}

func (u User) GetUserInfo() string{
	return fmt.Sprintf("Name: %s Age: %d", u.Name, u.Age)
}

func (u User) GetUserName() string{
	return u.Name
}

func (u User) GetUserAge() int{
	return u.Age
}

type User2 struct{
	Name string
	Age int
	Address string
}

func (u User2) GetUserInfo() string{
	return fmt.Sprintf("Name: %s Age: %d Address: %s", u.Name, u.Age, u.Address)
}

func (u User2) GetUserName() string{
	return u.Name + " " + "Doe"
}

func (u User2) GetUserAge() int{
	u.Age++
	return u.Age
}



func main(){
	user1 := User{
		Name: "Ruth",
		Age: 14,
	}
	fmt.Println(user1.GetUserInfo())
	fmt.Println(user1.GetUserName())
	fmt.Println(user1.GetUserAge())

	user2 := User2{
		Name: "Kathy",
		Age: 14,
		Address: "Heilshelm",
	}
	fmt.Println(user2.GetUserInfo())
	fmt.Println(user2.GetUserName())
	fmt.Println(user2.GetUserAge())

	var u UserInfo
	u= User{
		Name: "John",
		Age: 13,
	}
	fmt.Printf("%#v, %T", u, u)
}