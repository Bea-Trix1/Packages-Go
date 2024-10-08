package main

import "net/http"

func main() {
	mux := http.NewServeMux() // Cria um novo servidor mux ele serve para criar rotas e handlers

	http.HandleFunc("/", helloHandler) // Cria uma rota para a raiz do site
	http.Handle("/ola", &ola{"Oi"})    //	Cria uma rota para /ola e passa um handler para ela
	http.ListenAndServe(":8080", mux)  // Inicia o servidor na porta 8080
}

type ola struct {
	oi string
}

// ServeHTTP é um método que implementa a interface http.Handler
func (o *ola) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(o.oi)) // Escreve na resposta o texto que foi passado no handler
}

// helloHandler é um handler que escreve na resposta "Hello World"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
