package main

import (
	"fmt"
	"strings"
)

func ChisOuBola() string {
	var opcao string
	fmt.Println("\nEscolha uma opção: Letra X ou Letra O.")
	fmt.Scan(&opcao)
	if strings.ToUpper(opcao) != "X" && strings.ToUpper(opcao) != "O" {
		return ChisOuBola()
	}
	return strings.ToUpper(opcao)
}
