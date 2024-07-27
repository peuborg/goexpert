package main

import "net/http"

func main() {
	//http://localhost:8088/
	http.HandleFunc("/", BuscaCepHttp)
	http.ListenAndServe(":8088", nil)
}

func BuscaCepHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alô Mundo!"))
}
