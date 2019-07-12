package funcs

import "fmt"

func simple(a func(a int, b int) int) {
	fmt.Println("a(60, 7) is ", a(60, 7))
}

//RunHighOrderFunc runs high-ordered function simple()
func RunHighOrderFunc() {
	//simple func is high-ordered cause it apply func as incomong attribute
	f := func(a int, b int) int {
		return a + b
	}

	simple(f)
}
