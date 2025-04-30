package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Produto struct {
	ID    string
	Nome  string
	Preco float64
}

func main() {
	//sql.Open(<user>:<password>@tcp(<host>:<porta>)/<database>)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//O Ping simula uma conexão ao Banco de Dados para verificar se está tudo certo
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//insert
	/*produto := novoProduto("capa iphone", 19.90)
	err = inserirProduto(db, produto)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserirProduto - Procuto: %v, preço: %.2f\n", produto.Nome, produto.Preco)*/

	//update
	/*produto.Preco = 30.90
	produto.Nome = "capa iphone+"
	err = alterarProduto(db, produto)
	if err != nil {
		panic(err)
	}
	fmt.Printf("alterarProduto - Procuto: %v, preço: %.2f\n", produto.Nome, produto.Preco)*/

	//select one
	/*p, err := selecionarProduto(db, produto.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("selecionarProduto - Procuto: %v, possui o preço de %.2f\n", p.Nome, p.Preco)*/

	//select all
	produtos, err := selecionarTodosProdutos(db)
	if err != nil {
		panic(err)
	}
	for _, p := range produtos {
		fmt.Printf("selecionarTodosProdutos - ID: %v | Procuto: %v | Preço %.2f\n", p.ID, p.Nome, p.Preco)
	}
	//deletar
	/*err = deletarProduto(db, produto.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deletarProduto - Procuto: %v", produto.ID)*/

}

func novoProduto(nome string, preco float64) *Produto {
	return &Produto{
		ID:    uuid.New().String(),
		Nome:  nome,
		Preco: preco,
	}
}

func selecionarProduto(db *sql.DB, id string) (*Produto, error) {
	//Statemente prepare MySQL - Impede SQLInjection
	stmt, err := db.Prepare("select id, nome, preco from produtos where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var p Produto
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Nome, &p.Preco)
	if err != nil {
		panic(err)
	}
	return &p, nil
}

func selecionarTodosProdutos(db *sql.DB) ([]Produto, error) {
	//Statemente prepare MySQL - Impede SQLInjection
	rows, err := db.Query("select id, nome, preco from produtos order by nome")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var p Produto
		err = rows.Scan(&p.ID, &p.Nome, &p.Preco)
		if err != nil {
			panic(err)
		}
		produtos = append(produtos, p)
	}

	return produtos, nil
}

func inserirProduto(db *sql.DB, produto *Produto) error {
	//Statemente prepare MySQL - Impede SQLInjection
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

func alterarProduto(db *sql.DB, produto *Produto) error {
	//Statemente prepare MySQL - Impede SQLInjection
	stmt, err := db.Prepare("update produtos set nome = ?, preco = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.Nome, produto.Preco, produto.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func deletarProduto(db *sql.DB, id string) error {
	//Statemente prepare MySQL - Impede SQLInjection
	stmt, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return nil
}
