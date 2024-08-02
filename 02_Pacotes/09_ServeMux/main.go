package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Meu Blog Mux"})
	http.ListenAndServe(":8088", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Alô Peuborg!"))
	})
	http.ListenAndServe(":8089", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alô Mundo!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
