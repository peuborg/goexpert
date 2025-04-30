package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categoria struct {
	ID   int `gorm:"primaryKey;"`
	Nome string
}

type Item struct {
	ID          int `gorm:"primaryKey;"`
	Nome        string
	Preco       float64
	CategoriaID int
	Categoria   Categoria
	gorm.Model  //Cria as colunas datetime de ins upd del
}

func main() {
	//Open
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Create Table
	db.AutoMigrate(&Item{}, &Categoria{})

	//Insert Categoria
	/*categoria := Categoria{Nome: "Doces"}
	db.Create(&categoria)
	//Insert Belongs To
	db.Create(&Item{
		Nome:        "brigadeiro",
		Preco:       1,
		CategoriaID: categoria.ID,
	})*/

	//Select all
	var itens []Item
	//Select Belongs To (Item pertence à Categoria)
	db.Preload("Categoria").Find(&itens) //Preload traz os dados das Categorias, junto da busca dos itens
	for _, Item := range itens {
		fmt.Println(Item.Nome, Item.Categoria.Nome)
		fmt.Println(Item)
		fmt.Println()
	}
}
