package main

/*
func LimitesCima(coluna int) bool { //implementar em Jogadas
	return coluna > 0
}
func LimitesBaixo(coluna int) bool {
	return coluna < 3
}
func LimitesDireita(linha int) bool {
	return linha < 3
}

func LimitesEsquerda(linha int) bool {
	return linha > 0
}
*/
func Verify(slice [][]string) string {
	var resp string
	switch {
	//horizontal
	case (slice[0][0] == slice[0][1] && slice[0][0] == slice[0][2]) && slice[0][0] != " ":
		resp = "VENCEU"
	case (slice[1][0] == slice[1][1] && slice[1][0] == slice[1][2]) && slice[1][0] != " ":
		resp = "VENCEU"
	case (slice[2][0] == slice[2][1] && slice[2][0] == slice[2][2]) && slice[2][0] != " ":
		resp = "VENCEU"

	//diagonal
	case (slice[0][0] == slice[1][1] && slice[0][0] == slice[2][2]) && slice[0][0] != " ":
		resp = "VENCEU"
	case (slice[0][2] == slice[1][1] && slice[0][2] == slice[2][0]) && slice[0][2] != " ":
		resp = "VENCEU"

		//vertical
	case (slice[0][0] == slice[1][0] && slice[0][0] == slice[2][0]) && slice[1][0] != " ":
		resp = "VENCEU"
	case (slice[0][1] == slice[1][1] && slice[0][1] == slice[2][1]) && slice[0][1] != " ":
		resp = "VENCEU"
	case (slice[0][2] == slice[1][2] && slice[0][2] == slice[2][2]) && slice[0][2] != " ":
		resp = "VENCEU"

	//velha
	default:
		var testeVazio int
		for _, v := range slice {
			for _, value := range v {
				if value != " " {
					testeVazio++
				}
			}
		}
		if testeVazio == 9 {
			resp = "VELHA"
		}
	}
	return resp
}

/*if slice[0][0] == slice[0][1] && slice[0][0] == slice[0][2]) {
		resp = fmt.Sprintf("VENCEU. %s, %s, %s", slice[0][0], slice[0][1], slice[0][2])
	} else if (slice[0][0] == slice[1][0] && slice[0][0] == slice[2][0]) && (slice[0][0] != " ") {
		resp = fmt.Sprintf("VENCEU. %s, %s, %s", slice[0][0], slice[1][0], slice[2][0])
	}

	return resp
}


func Verifica(slice [][]string) string { // tentar verificação a partir da posição jogada
	switch {
	case slice[0][0] == slice[0][1] && slice[0][0] == slice[0][2]:
		return "VENCEU"
	case slice[1][0] == slice[1][1] && slice[1][0] == slice[1][2]:
		return "VENCEU"
	case slice[2][0] == slice[2][1] && slice[2][0] == slice[2][2]:
		return "VENCEU"
	case slice[0][0] == slice[1][1] && slice[0][0] == slice[2][2]:
		return "VENCEU"
	default:
		return "VELHA"
	}

}
*/
