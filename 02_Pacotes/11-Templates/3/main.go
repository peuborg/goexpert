package main

import (
	"html/template"
	"os"
)

type Disciplina struct {
	Nome         string
	CargaHoraria int
}

type Disciplinas []Disciplina

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Disciplinas{
		{"Banco de Dados", 60},
		{"Banco NoSql", 40},
		{"Desenvolvimento Front-End", 60},
		{"Desenvolvimento de APIs", 60},
	})
	if err != nil {
		panic(err)
	}
}
