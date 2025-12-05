package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter any number :", n)
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf(" %d * %d = %d\n ", n, i, n*i) ///multiplication table
	}
}
