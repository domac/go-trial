package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func main() {
	u := User{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)

	mv := User.TestValue
	mv(u)

	mp := (*User).TestPointer
	mp(&u)

	mp2 := (*User).TestValue
	mp2(&u)
}
