package main

import "fmt"

func ExibeJogo(slice [][]string) {
	const (
		coluna0 = iota
		coluna1
		coluna2
	)
	fmt.Println("\n")
	fmt.Println("  ", coluna0, coluna1, coluna2)
	fmt.Println("A", slice[0][:])
	fmt.Println("B", slice[1][:])
	fmt.Println("C", slice[2][:])
	fmt.Println("\n")
}
