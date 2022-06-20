package main

import "fmt"

var velha = make([][]string, 3)

func teste() {
	velha = [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	fmt.Println(velha)
	var cont int
	for _, v := range velha {
		for _, value := range v {
			if value == " " {
				cont++
			}
		}
	}
	fmt.Println(cont)
}
