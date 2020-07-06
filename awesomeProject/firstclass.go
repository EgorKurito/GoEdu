package main

import "fmt"

func main() {
	printer := func(in string) {
		fmt.Println("anon func out: ", in)
	}
	printer("test")

	type strFuncType func(string)

	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n\n", prefix, in)
		}
	}
	successLogger := prefixer("SUCCESS")
	successLogger("expected behavior")
}
