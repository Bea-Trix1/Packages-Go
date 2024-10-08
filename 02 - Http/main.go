package main

import (
	"fmt"
	"io"
	"net/http"
)

// Aqui estou fazendo uma requisição GET para o site do google e printando o conteúdo da página
func main() {
	req, err := http.Get("http://google.com") // O pacote http já faz a requisição GET, e assim eu aponto a URL que eu quero fazer a requisição
	if err != nil {
		panic(err)
	}

	resp, err := io.ReadAll(req.Body) // Aqui eu leio o corpo da resposta da requisição
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp)) // E printo o corpo da resposta
	req.Body.Close()          // É necessario fechar o corpo da resposta, para liberar os recursos e não ter vazamento de memória
}
