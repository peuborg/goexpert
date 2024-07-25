package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (vc ViaCep) Imprimir() string {
	var resposta string = "***CEP***"
	resposta += "\nCEP: " + vc.Cep
	resposta += "\nLogradouro: " + vc.Logradouro
	resposta += "\nComplemento: " + vc.Complemento
	resposta += "\nUnidade: " + vc.Unidade
	resposta += "\nBairro: " + vc.Bairro
	resposta += "\nLocalidade: " + vc.Localidade
	resposta += "\nUf: " + vc.Uf
	resposta += "\nIbge: " + vc.Ibge
	resposta += "\nGia: " + vc.Gia
	resposta += "\nDdd: " + vc.Ddd
	resposta += "\nSiafi: " + vc.Siafi

	return resposta
}

func main() {
	for _, url := range os.Args[1:] {
		//req, err := http.Get(url)			//comando: go run main.go http://viacep.com.br/ws/45604-116/json/
		req, err := http.Get("http://viacep.com.br/ws/" + url + "/json/") //comando: go run main.go 45604-116
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll((req.Body))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var cep ViaCep
		err = json.Unmarshal(res, &cep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println(cep)
		fmt.Println(cep.Imprimir())

		arq, err := os.Create("cep_" + url + ".txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer arq.Close()

		_, err = arq.WriteString(fmt.Sprintf(cep.Imprimir()))
	}
}
