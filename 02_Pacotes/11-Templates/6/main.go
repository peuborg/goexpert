package main

import (
	"html/template"
	"os"
	"strings"
)

type Disciplina struct {
	Nome         string
	CargaHoraria int
}

type Disciplinas []Disciplina

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Disciplinas{
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
