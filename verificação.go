package main

func Horizontal(slice [][]string) bool {
	return ((slice[0][0] == slice[0][1] && slice[0][0] == slice[0][2]) && slice[0][0] != " ") ||
		((slice[1][0] == slice[1][1] && slice[1][0] == slice[1][2]) && slice[1][0] != " ") ||
		((slice[2][0] == slice[2][1] && slice[2][0] == slice[2][2]) && slice[2][0] != " ")
}

func Diagonal(slice [][]string) bool {
	return ((slice[0][0] == slice[0][1] && slice[0][0] == slice[0][2]) && slice[0][0] != " ") ||
		((slice[0][2] == slice[1][1] && slice[0][2] == slice[2][0]) && slice[0][2] != " ")
}
func Vertical(slice [][]string) bool {
	return ((slice[0][0] == slice[1][0] && slice[0][0] == slice[2][0]) && slice[1][0] != " ") ||
		((slice[0][1] == slice[1][1] && slice[0][1] == slice[2][1]) && slice[0][1] != " ") ||
		((slice[0][2] == slice[1][2] && slice[0][2] == slice[2][2]) && slice[0][2] != " ")
}
func testeVazio(slice [][]string) bool {
	var testeVazio int
	for _, v := range slice {
		for _, value := range v {
			if value != " " {
				testeVazio++
			}
		}
	}
	return testeVazio == len(slice)*len(slice)
}

func Verify(slice [][]string) string {
	var resp string
	if Horizontal(slice) || Vertical(slice) || Diagonal(slice) {
		resp = "VENCEU"
	} else if !Horizontal(slice) && !Vertical(slice) && !Diagonal(slice) && testeVazio(slice) {
		resp = "VELHA"
	}
	return resp
}
