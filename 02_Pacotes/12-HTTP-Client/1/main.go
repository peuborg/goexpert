package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//criando o objeto http com limite de tempo para a resposta
	//c := http.Client{Timeout: time.Second}
	c := http.Client{Timeout: time.Microsecond}
	//executando a chamada http
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	/*	Defer é uma palavra reservada do Go que envia a função para uma lista de chamadas.
		Essas chamadas serão executadas depois que a função onde o defer se encontra terminar sua execução.	*/

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
