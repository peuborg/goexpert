package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pessoa struct {
	ID    int `gorm:"primaryKey;autoIncrement:true"`
	Nome  string
	Email string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//Create Table
	db.AutoMigrate(&Pessoa{})
	//Insert
	db.Create(&Pessoa{
		Nome:  "Fulano de Tal",
		Email: "fulano@email.com",
	})
	//Insert Many
	pessoas := []Pessoa{
		{Nome: "Beltrano", Email: "bel@email.com"},
		{Nome: "Ciclano", Email: "ciclano@email.com"},
		{Nome: "ze", Email: "ze@email.com"},
	}
	db.Create(pessoas)
}
