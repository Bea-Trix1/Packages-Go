package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	json := bytes.NewBuffer([]byte(`{"nome": "Beatriz"}`))             // crio um buffer de bytes com o JSON que eu quero enviar
	resp, err := c.Post("http://google.com", "application/json", json) // faço uma requisição POST para o site do google, passando o JSON que eu criei
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil) // Ele pega os dados, e escolhemos a onde jogamos esses dados e onde vou jogar os dados

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

}
