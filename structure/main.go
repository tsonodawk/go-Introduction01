package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
	// x, y int
}

func UpdateUser(user User) {
	user.Name = "update"
	user.Age = 1000
	// fmt.Println(user)
}

func UpdateUser2(user *User) {
	user.Name = "update"
	user.Age = 1000
	// fmt.Println(user)
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {

	// struct1()

	// strunctMethod()

	// strunctEmbed()

	strunctConstract()
}

func strunctConstract() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1)
	fmt.Println(*user1)
}

func strunctEmbed() {
	t := T{User: User{Name: "user1"}}
	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	t.User.SetName2("SetA")
	t.User.SayName()

}

func strunctMethod() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("A")
	user1.SayName()

}

func struct1() {
	var user1 User
	fmt.Println(user1)

	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{"user2", 20}
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	fmt.Println(user1)
	UpdateUser2(user8)
	fmt.Println(user8)
}
