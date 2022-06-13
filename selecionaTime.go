package main

import (
	"fmt"
	"strings"
)

func Select() string {
	var option string
	fmt.Scanf("%s", &option)
	return strings.ToUpper(option)
}
func Jogadores() {
	jogador1 = Select()
	if jogador1 == "X" {
		jogador2 = "O"
	} else {
		jogador2 = "X"
	}
}
