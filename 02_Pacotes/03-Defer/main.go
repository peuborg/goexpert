package main

import (
	"io"
	"net/http"
)

/*
Defer é uma palavra reservada do Go que envia a função para uma lista de chamadas.
Essas chamadas serão executadas depois que a função onde o defer se encontra terminar sua execução.
*/

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll((req.Body))
	if err != nil {
		panic(err)
	}

	println(string(res))
	req.Body.Close()
}
