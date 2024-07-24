package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	//ao final da execução é necessário fechar o req.body
	defer req.Body.Close() //o comando defer vai fazer com que a instrução seja a última executada

	res, err := io.ReadAll((req.Body))
	if err != nil {
		panic(err)
	}

	println(string(res))

}
