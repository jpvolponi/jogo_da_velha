package main

import (
	"fmt"
)

//jogadores
var jogador1, jogador2, linha string
var coluna, row int

var tralha = [][]string{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

func JogoDaVelha(slice [][]string) {
	fmt.Println("Escolha uma opção: X ou O")
	Jogadores(&jogador1, &jogador2)
	fmt.Printf("jogador1: %s\njogador2: %s\n", jogador1, jogador2)
	ExibeJogo(tralha)
	Jogadas(slice, jogador1)
	Jogadas(slice, jogador2)

}
