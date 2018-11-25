package main

import "fmt"

func main() {
	fmt.Println("Hello, Arvind")
	var count = 10
	for index := 0; index <= count; index++ {
		if index%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	// printMe("Hi, there!", "I'm Arvind Jaiswal", "I'm a Software Developer")
	printMe(y)
	foo()
	printMe(y)
}

func printMe(a ...interface{}) {
	// n, error := fmt.Println(a...)
	n, _ := fmt.Println(a...)
	fmt.Println("n = ", n)
}

func foo() {
	var x = 43
	y = 100
	fmt.Println("Down ", x)
}

var y = 1000
