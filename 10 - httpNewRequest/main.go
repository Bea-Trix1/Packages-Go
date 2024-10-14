package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil) // newRequest é uma função que cria uma nova requisição, passando o método, a URL e o corpo da requisição
	if err != nil {
		panic(err)
	}

	req.Header.Set("Aceppt", "application/json") // Aqui eu seto o cabeçalho da requisição, passando o tipo de conteúdo que eu quero enviar
	resp, err := c.Do(req)                       // Do é um método que faz a requisição, passando a requisição que eu criei
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
