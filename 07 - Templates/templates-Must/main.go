package main

import (
	"html/template" // html/template é um pacote que permite a criação de templates de forma segura, evitando ataques de injeção de código
	"os"            // os é um pacote que fornece uma plataforma independente de acesso ao sistema operacional
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} | Carga Horária: {{.CargaHoraria}} horas")) // Cria um template em uma única linha
	err := t.Execute(os.Stdout, curso)                                                                                   // Executa o template passando os dados do curso
	if err != nil {
		panic(err)
	}
}
