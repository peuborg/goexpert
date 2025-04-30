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

type Objeto struct {
	ID           int `gorm:"primaryKey;"`
	Nome         string
	Preco        float64
	CategoriaID  int
	Categoria    Categoria
	NumeroSerial NumeroSerial
	gorm.Model   //Cria as colunas datetime de ins upd del
}

type NumeroSerial struct {
	ID       int `gorm:"primaryKey;"`
	Numero   string
	ObjetoID int
}

func main() {
	//Open
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Create Table
	db.AutoMigrate(&Objeto{}, &Categoria{}, &NumeroSerial{})

	//Insert Categoria
	categoria := Categoria{Nome: "coisas em garrafa"}
	db.Create(&categoria)
	//Insert Belongs To
	objeto := Objeto{
		Nome:        "diversos em garrafa",
		Preco:       300,
		CategoriaID: categoria.ID,
	}
	db.Create(&objeto)
	//Insert Has One
	db.Create(&NumeroSerial{
		Numero:   "456",
		ObjetoID: objeto.ID,
	})

	//Select all
	var objetos []Objeto
	//Select Belongs To e Has One
	//Objeto Pertence a Categoria e Tem Um NumeroSerial
	db.Preload("Categoria").Preload("NumeroSerial").Find(&objetos)
	for _, Objeto := range objetos {
		fmt.Println(Objeto)
		fmt.Println(Objeto.Nome, Objeto.Categoria.Nome, Objeto.NumeroSerial.Numero)
	}
}
