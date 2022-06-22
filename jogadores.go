package main

func Jogadores(p1, p2 *string) (string, string) {
	switch ChisOuBola() {
	case "X":
		*p1 = "X"
		*p2 = "O"
	case "O":
		*p1 = "O"
		*p2 = "X"
	}
	return *p1, *p2
}
