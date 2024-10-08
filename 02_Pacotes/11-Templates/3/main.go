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
	err := t.Execute(os.Stdout, Disciplinas{ //Neste caso, a saída será no console (os.Stdout)
		{"Banco de Dados", 60},
		{"Banco NoSql", 40},
		{"Desenvolvimento Front-End", 60},
		{"Desenvolvimento de APIs", 60},
		{"Programação Estruturada", 10},
		{"Programação Orientada a Objetos", 10},
	})
	if err != nil {
		panic(err)
	}
}
