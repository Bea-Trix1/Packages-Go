package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second} // Aqui eu crio um cliente http, que é um objeto que me permite fazer requisições http
	resp, err := c.Get("http://google.com")                   // Aqui eu faço uma requisição http para o google
	if err != nil {
		panic(err) // Se houver um erro, eu paro a execução do programa
	}

	defer resp.Body.Close() // Aqui eu fecho o corpo da resposta, para que não haja vazamento de memória

	body, err := io.ReadAll(resp.Body) // Aqui eu leio o corpo da resposta
	if err != nil {
		panic(err)
	}

	println(string(body)) // Aqui eu imprimo o corpo da resposta
}
