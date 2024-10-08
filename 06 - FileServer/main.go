package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()                           // Cria um novo servidor mux ele serve para criar rotas e handlers
	fileServer := http.FileServer(http.Dir("./public")) // Cria um servidor de arquivos que serve os arquivos da pasta public
	mux.Handle("/", fileServer)                         // Cria uma rota para a raiz do site
	log.Fatal(http.ListenAndServe(":8080", mux))        // Inicia o servidor na porta 8080
}
