package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pessoa struct {
	ID    int `gorm:"primaryKey;autoIncrement:true"`
	Nome  string
	Email string
}

func main() {
	//Open
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Create Table
	db.AutoMigrate(&Pessoa{})

	//Insert
	/*db.Create(&Pessoa{
		Nome:  "Fulano de Tal",
		Email: "fulano@email.com",
	})*/

	//Insert Many
	pessoas := []Pessoa{
		{Nome: "Bel", Email: "bel@email.com"},
		{Nome: "Cicl", Email: "cicl@email.com"},
		{Nome: "ze", Email: "ze@email.com"},
	}
	db.Create(pessoas)

	//Select one
	/*var pessoa Pessoa
	//Seleciona o primeiro registro com o código do parâmetro
	db.First(&pessoa, 1)
	fmt.Println(pessoa)
	//Seleciona o primeiro registro que satisfaz a condição
	db.First(&pessoa, "nome = ?", "ze")
	fmt.Println(pessoa)*/

	//Select all
	/*db.Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}*/

	//Select os 5 primeiros registros
	/*db.Limit(5).Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}*/

	//Select os 5 primeiros registros da página 2
	/*db.Limit(5).Offset(5).Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}*/

	//WHERE
	/*db.Where("id > ?", 5).Find(&pessoas)
	for _, pessoa := range pessoas {,.;
		fmt.Println(pessoa)
	}*/
	db.Where("nome LIKE ?", "%Ciclano%").Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}
}
