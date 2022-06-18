package main

import (
	"fmt"
	"strings"
)

func LimitesCima(i int) bool { //implementar em Jogadas
	return i > 0
}
func LimitesBaixo(i, tamanhoVelha int) bool {
	return i < tamanhoVelha-1
}
func LimitesDireita(j, tamanhoVelha int) bool {
	return j < tamanhoVelha-1
}

func LimitesEsquerda(j int) bool {
	return j < 2
}

func ExibeJogo(slice [][]string) {
	fmt.Println("\n")
	fmt.Println("   0 1 2")
	fmt.Println("A", slice[0][:])
	fmt.Println("B", slice[1][:])
	fmt.Println("C", slice[2][:])
	fmt.Println("\n")
}
func ChisOuBola() string {
	var option string
	fmt.Scan(&option)
	return strings.ToUpper(option)
}

func Jogadores(p1, p2 *string) (string, string) {
	switch ChisOuBola() {
	case "X":
		*p1 = "X"
		*p2 = "O"
	case "O":
		*p1 = "O"
		*p2 = "X"
	}
	return *p1, *p2
}

func Jogadas(slice [][]string, player string) {
	var coluna, linha int
	var row string
	fmt.Println("\n\nEscolha a linha da jogada:")
	fmt.Scan(&row)
	fmt.Println("\n\nEscolha a coluna da jogada:")
	fmt.Scan(&coluna)
	switch strings.ToUpper(row) {
	case "A":
		linha = 0
	case "B":
		linha = 1
	case "C":
		linha = 2
	}
	slice[linha][coluna] = player
	ExibeJogo(slice)
}
