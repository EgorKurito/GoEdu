package main

import (
	"Go/packTest/visibility/person"
	"fmt"
)

func main()  {
	p := person.NewPerson(1, "EgorKurito", "secret")

	// fmt.Printf("main.PrintPerson: %+v\n", p.secret)

	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
}
