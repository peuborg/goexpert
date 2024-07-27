package main

import (
	"encoding/json"
	"io"
	"net/http"
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

func main() {
	//http://localhost:8088/
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8088", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	/*res, error = json.Marshal(cep)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)*/
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCep, error) {
	req, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/") //comando: go run main.go 45604-116
	if error != nil {
		return nil, error
	}
	defer req.Body.Close()

	res, error := io.ReadAll((req.Body))
	if error != nil {
		return nil, error
	}

	var c ViaCep
	error = json.Unmarshal(res, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
