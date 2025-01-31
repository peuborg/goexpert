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
	ID           int `gorm:"primaryKey;"`
	Nome         string
	Preco        float64
	CategoriaID  int
	Categoria    Categoria
	NumeroSerial NumeroSerial
	gorm.Model   //Cria as colunas datetime de ins upd del
}

type NumeroSerial struct {
	ID     int `gorm:"primaryKey;"`
	Numero string
	ItemID int
}

func main() {
	//Open
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Create Table
	db.AutoMigrate(&Item{}, &Categoria{}, &NumeroSerial{})

	//Insert Categoria
	categoria := Categoria{Nome: "calçados"}
	db.Create(&categoria)
	//Insert Belongs To
	item := Item{
		Nome:        "tenis adidas",
		Preco:       1200,
		CategoriaID: categoria.ID,
	}
	db.Create(&item)
	//Insert Has One
	db.Create(&NumeroSerial{
		Numero: "1234567890",
		ItemID: item.ID,
	})

	//Select all
	var itens []Item
	//Select Belongs To e Has One
	db.Preload("Categoria").Preload("NumeroSerial").Find(&itens)
	for _, Item := range itens {
		fmt.Println(Item)
		fmt.Println(Item.Nome, Item.Categoria.Nome, Item.NumeroSerial.Numero)
	}
}
