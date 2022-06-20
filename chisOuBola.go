package main

import (
	"fmt"
	"strings"
)

func ChisOuBola() string {
	var opcao string
	//var opcaoFinal string
	fmt.Println("\nEscolha uma opção: Letra X ou Letra O.")
	fmt.Scan(&opcao)

	/*
		if strings.ToUpper(opcao) != "X" && strings.ToUpper(opcao) != "O" {
			ChisOuBola()
		} else {
			opcaoFinal = strings.ToUpper(opcao)
			fmt.Println(opcaoFinal)
		}
		return strings.ToUpper(opcao)
	*/
	return strings.ToUpper(opcao)
}
