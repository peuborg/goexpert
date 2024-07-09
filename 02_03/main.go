package main

const ip = "127.0.0.1"

type login string

var (
	conectou bool    = true
	tempo    int     = 10
	resposta string  = "sucesso"
	lucro    float64 = 500.23
	usuario  login   = "pedro"
)

func main() {
	nome := "Peuborg"
	println(ip, conectou, tempo, resposta, lucro, nome, usuario)
}
