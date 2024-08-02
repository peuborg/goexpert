package main

import (
	"html/template"
	"os"
)

type Disciplina struct {
	Nome         string
	CargaHoraria int
}

func main() {
	semestre2023_2 := Disciplina{"Banco de Dados", 60}
	t := template.Must(template.New("DisciplinaTemplate").Parse("Disciplina: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, semestre2023_2)
	if err != nil {
		panic(err)
	}
}
