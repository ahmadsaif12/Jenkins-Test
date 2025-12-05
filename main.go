package main

import "fmt"

func main() {
	// print even numbers
	var n int
	fmt.Println(" Enter any number :", n)
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Printf("%d is even number \n", n)
	} else {
		fmt.Printf("%d is odd number\n", n)
	}
}
