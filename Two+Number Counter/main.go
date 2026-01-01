package main

import "fmt"

func Slozh(number1 int, number2 int) int {
	rezultat := number1 + number2
	return rezultat
}

func main() {
	var a int
	var b int

	fmt.Println("Enter two numbers (like this -> 123 123)")
	fmt.Scan(&a)
	fmt.Scan(&b)

	itog := Slozh(a, b)
	fmt.Println("Result:", itog)
}
