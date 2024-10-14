package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Um contexto é um objeto que carrega informações sobre o ambiente em que a aplicação está rodando
// Ele é utilizado para cancelar requisições, definir timeouts, entre outras coisas
// O contexto é passado como parâmetro para várias funções, como o http.NewRequestWithContext
// O http.NewRequestWithContext é uma função que cria uma nova requisição, passando o contexto, o método, a URL e o corpo da requisição
// Uma analogia para o contexto é o cancelamento de uma requisição, se você está fazendo uma requisição e ela está demorando muito, você pode cancelar essa requisição
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
