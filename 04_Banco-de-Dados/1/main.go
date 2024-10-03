package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Produto struct {
	ID    string
	Nome  string
	Preco float64
}

func novoProduto(nome string, preco float64) *Produto {
	return &Produto{
		ID:    uuid.New().String(),
		Nome:  nome,
		Preco: preco,
	}
}

func inserirProduto(db *sql.DB, produto *Produto) error {
	//Statemente prepare MySQL
	stmt, err := db.Prepare("insert into produtos(id, nome, preco) values(?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.ID, produto.Nome, produto.Preco)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	//sql.Open(user:password@tcp(localhost:3306)/database)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	produto := novoProduto("iphone", 2199.90)
	err = inserirProduto(db, produto)
	if err != nil {
		panic(err)
	}
}
