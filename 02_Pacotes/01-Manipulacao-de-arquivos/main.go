package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//***Criando um arquivo***
	f, err := os.Create("arq.txt")
	if err != nil {
		panic(err)
	}

	//Escrevendo no arquivo
	//tamanho, err := f.WriteString("alô Mundo")
	tamanho, err := f.Write([]byte("alo Mundao"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	//Fechando o arquivo
	f.Close()

	//***Leitura de um arquivo***
	arq, err := os.ReadFile("arq.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arq))

	//***Leitura de pouco em pouco***
	arq2, err := os.Open("arq.txt")
	if err != nil {
		panic(err)
	}
	meuReader := bufio.NewReader(arq2)
	meuBuffer := make([]byte, 4)
	for {
		n, err := meuReader.Read(meuBuffer)
		if err != nil {
			break
		}
		fmt.Println(string(meuBuffer[:n]))
	}

	//***Removendo um arquivo***
	err = os.Remove("arq.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Arquivo apagado com sucesso!")
	}
}
