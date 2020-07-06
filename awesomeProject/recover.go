package main

import (
	"fmt"
)

func deferTest()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happend: ", err)
		}
	}()
	fmt.Println("Some useful work")
	panic("something bad happend")
	return
}

func main()  {
	deferTest()
	return
}