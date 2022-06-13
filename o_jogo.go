package main

import (
	"fmt"
	"strings"
)

//jogadores
var jogador1, jogador2, linha string
var coluna, row int

var tralha = [][]string{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

func main() {
	fmt.Println("Escolha uma opção: X ou O")
	Jogadores()
	fmt.Printf("jogador1: %s\njogador2: %s\n", jogador1, jogador2)
	fmt.Println("   0 1 2")
	fmt.Println("A", tralha[0][:])
	fmt.Println("B", tralha[1][:])
	fmt.Println("C", tralha[2][:])
	fmt.Println("\n\nEscolha a linha e a coluna da jogada:")
	fmt.Scanln("%s", &linha)
	fmt.Println("\n\nEscolha a linha e a coluna da jogada:")
	fmt.Scanln("%d", &coluna)

	//fmt.Println("\n\nEscolha a coluna da jogada:")

	switch strings.ToUpper(linha) {
	case "A":
		row = 0
	case "B":
		row = 1
	case "C":
		row = 2
	}

	tralha[row][coluna] = jogador1
	fmt.Println("\n   0 1 2")
	fmt.Println("A", tralha[0][:])
	fmt.Println("B", tralha[1][:])
	fmt.Println("C", tralha[2][:])
}
