package main

import (
	"html/template"
	"net/http"
)

type Disciplina struct {
	Nome         string
	CargaHoraria int
}

type Disciplinas []Disciplina

func main() {
	//http://localhost:8088/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Disciplinas{ //Neste caso, a saída será no webserver w
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
	})
	http.ListenAndServe(":8088", nil)

}
