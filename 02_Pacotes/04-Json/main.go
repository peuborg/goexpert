package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	//***Serializando_1
	res, err := json.Marshal(conta) //.Marshal converte/serializa a struct em JSON
	if err != nil {
		panic(err)
	}
	println(string(res))

	//***Serializando_2
	err = json.NewEncoder(os.Stdout).Encode(conta) //.Encoder já converte/serializa devolvendo para algum lugar, no caso o terminal
	if err != nil {
		panic(err)
	}

	//***Deserializando_1
	jsonPuro := []byte(`{"numero":2,"saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX) //.Unmarshal converte/deserializa o JSON em uma struct
	if err != nil {
		panic(err)
	}
	fmt.Printf("ContaX: %d Saldo: %d\n", contaX.Numero, contaX.Saldo)

	//***Deserializando_2
	jsonPuro2 := []byte(`{"n":3,"s":300}`)
	var contaY Conta
	err = json.Unmarshal(jsonPuro2, &contaY) //.Unmarshal converte/deserializa o JSON em uma struct
	if err != nil {
		panic(err)
	}
	fmt.Printf("ContaY: %d Saldo: %d\n", contaY.Numero, contaY.Saldo)
}
