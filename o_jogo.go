package main

import (
	"fmt"
)

//jogadores
var jogador1, jogador2, linha string

type jogador struct {
	titulo string
	time   string
}

var tralha [][]string

func JogoDaVelha(slice [][]string) {
	var result, vencedor string
	jogador1 := jogador{titulo: "Jogador 1"}
	jogador2 := jogador{titulo: "Jogador 2"}

	fmt.Println("Escolha uma opção: X ou O")
	Jogadores(&jogador1.time, &jogador2.time)
	fmt.Printf("jogador1: %s\njogador2: %s\n", jogador1, jogador2)
	ExibeJogo(slice)

	for result != "VELHA" && result != "VENCEU" {
		Jogadas(slice, jogador1)
		result = Verify(slice)
		vencedor = jogador1.titulo
		if result == "VELHA" || result == "VENCEU" {
			continue
		}
		Jogadas(slice, jogador2)
		result = Verify(slice)
		vencedor = jogador2.titulo
	}
	if result == "VELHA" {
		vencedor = "VELHA"
		fmt.Printf("Fim de jogo.\nNão houve vencedores nesta partida.\nO resultado do jogo é %s.", vencedor)
	} else {
		fmt.Printf("Fim de jogo.\nO %s venceu a partida.\nParabéns!!!!!", vencedor)
	}

}
