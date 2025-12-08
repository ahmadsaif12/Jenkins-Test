package main

import "fmt"

// func adding two numbers
func Add(a int, b int) int {
	return a + b
}

func main() {

	result := Add(8, 4)
	fmt.Printf(" the additon of two number is %d", result)
	// print even numbers
	var n int
	fmt.Println(" Enter any number :")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Printf("%d is even number \n", n)
	} else {
		fmt.Printf("%d is odd number\n", n)
	}
	fmt.Println("hey changes")
}
