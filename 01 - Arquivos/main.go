package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Criando um arquivo com o pacote OS
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escrevendo no arquivo uma string
	// tamanho, err := f.WriteString("Olá, mundo!\n")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Tamanho: %d bytes\n", tamanho)

	// Se eu não souber o que eu vou gravar no arquivo, posso usar Write
	// Write recebe um slice de bytes
	bytes, err := f.Write([]byte("Escrevendo dados no meu arquivo\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes: %d\n", bytes)
	f.Close()

	// Lendo arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Leitura linha a linha abrindo o arquivo
	arq, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Ele vai ler 16 bytes por vez atraves do pacote io
	// e vai armazenar no buffer o tamanho de 16 bytes
	reader := bufio.NewReader(arq)
	buffer := make([]byte, 3)

	// Loop onde o arquivo é lido, vai pegar os dados e atribuir ao buffer
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// removendo arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
