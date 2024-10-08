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
	tmp := template.New("CursoTemplate")                                            // Cria um novo template com o nome CursoTemplate
	tmp, _ = tmp.Parse("Curso: {{.Nome}} | Carga Horária: {{.CargaHoraria}} horas") // Adiciona um template ao template criado
	err := tmp.Execute(os.Stdout, curso)                                            // Executa o template passando os dados do curso
	if err != nil {
		panic(err)
	}
}
