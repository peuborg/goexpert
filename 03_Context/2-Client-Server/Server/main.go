package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		//Executa o processamente e dá um retorno após 5 segundos
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso!"))
	case <-ctx.Done():
		//Cancela o processamento caso o Client cancele a requisição
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Requeste cancelada pelo cliente", http.StatusRequestTimeout)
	}

}
