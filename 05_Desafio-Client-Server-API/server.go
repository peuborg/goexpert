package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

// Estrutura para mapear a resposta da API externa
type CotacaoResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", cotacaoHandler)

	log.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	// Contexto com timeout de 200ms para chamada da API
	ctxAPI, cancelAPI := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelAPI()

	req, err := http.NewRequestWithContext(ctxAPI, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Erro ao criar requisição", http.StatusInternalServerError)
		log.Println("Erro ao criar requisição:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Erro ao obter cotação", http.StatusRequestTimeout)
		log.Println("Erro ao obter cotação:", err)
		return
	}
	defer resp.Body.Close()

	var cotacao CotacaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		http.Error(w, "Erro ao decodificar resposta", http.StatusInternalServerError)
		log.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Contexto com timeout de 10ms para salvar no banco
	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	if err := salvarCotacao(ctxDB, cotacao.USDBRL.Bid); err != nil {
		log.Println("Erro ao salvar cotação no banco:", err)
	}

	// Retorna somente o valor do bid
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": cotacao.USDBRL.Bid})
}

func salvarCotacao(ctx context.Context, bid string) error {
	db, err := sql.Open("sqlite", "./cotacoes.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, "INSERT INTO cotacoes (bid) VALUES (?)", bid)
	return err
}
