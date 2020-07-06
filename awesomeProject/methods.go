package main

import (
	"fmt"
)

type Person struct {
	Id		int
	Name	string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id		int
	Name	string
	Person
}

func main()  {
	//pers := &Person{1, "Egor"}
	pers := new(Person)
	pers.SetName("Egor Kurito")
	//(&pers).SetName("Egor Kurito")
	fmt.Printf("Updated person: %#v\n", pers)

	var acc Account = Account {

	}
}