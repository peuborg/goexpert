package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Estrutura para extrair apenas o campo "bid"
type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	// Cria contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Cria requisição com o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro ao fazer requisição:", err)
	}
	defer resp.Body.Close()

	// Lê e decodifica o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao ler resposta:", err)
	}

	var cotacao Cotacao
	if err := json.Unmarshal(body, &cotacao); err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	// Formata e escreve no arquivo
	conteudo := fmt.Sprintf("Dólar: %s\n", cotacao.Bid)
	err = os.WriteFile("cotacao.txt", []byte(conteudo), 0644)
	if err != nil {
		log.Fatal("Erro ao escrever no arquivo:", err)
	}

	fmt.Println("Cotação salva com sucesso:", conteudo)
}
