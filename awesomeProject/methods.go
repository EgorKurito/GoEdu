package main

import (
	"fmt"
)

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int)  {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int  {
	return len(*sl)
}

func main() {
	//pers := &Person{1, "Egor"}
	pers := new(Person)
	pers.SetName("Egor Kurito")
	//(&pers).SetName("Egor Kurito")
	fmt.Printf("Updated person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "EKurito",
		Person: Person{
			Id:   2,
			Name: "Egor Kurito",
		},
	}
	fmt.Printf("%#v \n", acc)

	acc.SetName("egor.kurito")
	fmt.Printf("%#v \n", acc)

	acc.Person.SetName("egor.kurito")
	fmt.Printf("%#v \n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(6)
	fmt.Println(sl.Count(), sl)
}
