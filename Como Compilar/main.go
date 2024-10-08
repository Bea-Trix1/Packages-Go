package main

// Para compilar o arquivo main.go, execute:
// go build main.go

// Isso vai gerar um arquivo executável com o nome do arquivo sem a extensão .go
// No caso, main

// Para executar o arquivo, execute:
// ./main

// Isso permite que você compile e execute o arquivo sem precisar de um ambiente Go instalado

func main() {
	println(Soma(3, 3))
}

func Soma(a, b int) int {
	return a + b
}
