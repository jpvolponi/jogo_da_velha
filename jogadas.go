package main

import (
	"fmt"
	"strings"
)

func Jogadas(slice [][]string, player jogador) {
	var coluna, linha int
	var row string
	fmt.Printf("%s(%s), baseado no quadro acima, escolha a posição de sua jogada.", player.titulo, player.time)
	fmt.Println("\n\nInforme a LINHA e a COLUNA da jogada:")
	fmt.Scan(&row, &coluna)
	switch strings.ToUpper(row) {
	case "A":
		linha = 0
	case "B":
		linha = 1
	case "C":
		linha = 2
	}
	if slice[linha][coluna] == " " {
		slice[linha][coluna] = player.time
	} else {
		fmt.Printf("\nPosição indisponível!\n\n")
		Jogadas(slice, player)
	}

	ExibeJogo(slice)
	fmt.Printf("Última jogada: %s, %s.\n\n", player.time, player.titulo)
	fmt.Println("x-o-x-o-x-o-x-o-x-o-x-o-x-o-x\n\n")

}
