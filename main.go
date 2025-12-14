package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	calculate(10, 20)
}

func calculate(a int, b int)  {
	fmt.Println("Addition of", a, "and", b, "is", adding(a, b))
	fmt.Println("Subtraction of", a, "and", b, "is", subtracting(a, b))
}

func adding(a int, b int) int {
	return a + b
}

func subtracting(a int, b int) int {
	return a - b
}