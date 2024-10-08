package main

// https://pkg.go.dev/html/template -> Documentação do pacote html/template
import (
	"html/template" // html/template é um pacote que permite a criação de templates de forma segura, evitando ataques de injeção de código
	"net/http"
	// os é um pacote que fornece uma plataforma independente de acesso ao sistema operacional
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Cria um servidor web e traz o template
		t := template.Must(template.New("template.html").ParseFiles("template.html")) // Cria um template a partir de um arquivo - template.html
		err := t.Execute(w, Cursos{                                                   // Executa o template passando os dados do curso
			{"Go", 40},
			{"Java", 60},
			{"Python", 45},
			{"C++", 55},
			{"C#", 50},
		}) // Executa o template passando os dados do curso
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)

}
