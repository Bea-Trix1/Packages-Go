package main

import (
	"fmt"
)

// defer vai executar a função após o fim da execução da função main
// defer é muito útil para fechar arquivos, conexões, etc.
// é uma forma de garantir que o recurso será fechado após o uso
func main() {
	defer fmt.Println("primeira linha")
	fmt.Println("segunda linha")
	fmt.Println("terceira linha")

}
