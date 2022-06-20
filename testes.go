package main

var tralhaTeste = [][]string{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

/*
func Jogadas(slice [][]string, player string) {
	var result string
	for strings.ToUpper(result) != "VELHA" || strings.ToUpper(result) != "VENCEU" {
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
		result = Verifica(slice)
	}
}

func main() {
	/*var j1, j2 string

		ExibeJogo(tralhaTeste)
		tralhaTeste[0][2] = "X"
		tralhaTeste[2][2] = "O"
		ExibeJogo(tralhaTeste)

	fmt.Println(Jogadores(&j1, &j2))
	fmt.Println("JOGADOR 1:", j1)
	fmt.Println("JOGADOR 2:", j2)



}
*/
