package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second} // crio um cliente http, com timeout de 1 segundo
	resp, err := c.Get("http://google.com")                   // faço uma requisição http para o google
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // fecho o corpo da resposta

	body, err := io.ReadAll(resp.Body) // leio o corpo da resposta com io.ReadAll
	if err != nil {
		panic(err)
	}

	println(string(body))

}
