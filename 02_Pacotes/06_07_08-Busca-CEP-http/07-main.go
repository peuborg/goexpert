package main

import "net/http"

func main() {
	//http://localhost:8088/
	http.HandleFunc("/", BuscaCepHttp2)
	http.ListenAndServe(":8088", nil)
}

func BuscaCepHttp2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alô Mundo!"))
}
