package main

import "fmt"

func ExibeJogo(slice [][]string) {
	fmt.Println("\n")
	fmt.Println("   0 1 2")
	fmt.Println("A", slice[0][:])
	fmt.Println("B", slice[1][:])
	fmt.Println("C", slice[2][:])
	fmt.Println("\n")
}
