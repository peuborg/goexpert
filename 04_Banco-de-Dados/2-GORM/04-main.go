package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categoria struct {
	ID      int `gorm:"primaryKey;"`
	Nome    string
	Objetos []Objeto
}

type Objeto struct {
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
	db.AutoMigrate(&Objeto{}, &Categoria{})

	//Insert Categoria
	/*categoria := Categoria{Nome: "eletrodomésticos"}
	db.Create(&categoria)
	//Insert Belongs To
	objeto := Objeto{
		Nome:        "Geladeira",
		Preco:       5000,
		CategoriaID: categoria.ID,
	}
	db.Create(&objeto)*/

	//Has Many - Uma categoria pode ter vários produtos
	var categorias []Categoria
	err = db.Model(&Categoria{}).Preload("Objetos").Find(&categorias).Error
	if err != nil {
		panic(err)
	}

	//Exibindo todos os produtos de todas as categorias
	for _, categoria := range categorias {
		fmt.Println(categoria.Nome, ":")
		for _, objeto := range categoria.Objetos {
			fmt.Println("Objeto: ", objeto.Nome, categoria.Nome)
		}
	}
}
