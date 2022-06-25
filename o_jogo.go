package main

import (
	"fmt"
)

type jogador struct {
	titulo string
	time   string
}

func JogoDaVelha() {
	var tralha = [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	var result, vencedor string
	jogador1 := jogador{titulo: "Jogador 1"}
	jogador2 := jogador{titulo: "Jogador 2"}
	fmt.Printf("\n\n")
	fmt.Printf(`
	x-o-x-o-x-o-x-o-x-o-x-o-x-o-x
	o       JOGO DA VELHA       o
	x-o-x-o-x-o-x-o-x-o-x-o-x-o-x`)
	fmt.Println("")
	Jogadores(&jogador1.time, &jogador2.time)
	fmt.Printf("\nJogador n° 1: %s\nJogador n° 2: %s\n", jogador1.time, jogador2.time)
	ExibeJogo(tralha)

	for result != "VELHA" && result != "VENCEU" {
		Jogadas(tralha, jogador1)
		result = Resultado(tralha)
		vencedor = jogador1.titulo
		if result == "VELHA" || result == "VENCEU" {
			continue
		}
		Jogadas(tralha, jogador2)
		result = Resultado(tralha)
		vencedor = jogador2.titulo
	}
	if result == "VELHA" {
		vencedor = "VELHA"
		fmt.Printf("Fim de jogo.\nNão houve vencedores nesta partida.\nO resultado do jogo é %s.", vencedor)
	} else {
		fmt.Printf("Fim de jogo.\nO %s venceu a partida.\nParabéns!!!!!", vencedor)
	}

}
