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
	tmp := template.New("DisciplinaTemplate")
	tmp, _ = tmp.Parse("Disciplina: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, semestre2023_2)
	if err != nil {
		panic(err)
	}
}
