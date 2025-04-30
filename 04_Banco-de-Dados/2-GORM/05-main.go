package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Tipo struct {
	ID       int `gorm:"primaryKey;"`
	Nome     string
	Servicos []Servico `gorm:"many2many:servicos_tipos;"`
}

type Servico struct {
	ID         int `gorm:"primaryKey;"`
	Nome       string
	Preco      float64
	Tipos      []Tipo `gorm:"many2many:servicos_tipos;"`
	gorm.Model        //Cria as colunas datetime de ins upd del
}

func main() {
	//Open
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Create Table
	db.AutoMigrate(&Servico{}, &Tipo{})

	//Insert Tipo
	/*tipo := Tipo{Nome: "Acabamento"}
	db.Create(&tipo)
	tipo2 := Tipo{Nome: "Reforma"}
	db.Create(&tipo2)

	//Insert Servico
	servico := Servico{
		Nome:  "Gesso",
		Preco: 500,
		Tipos: []Tipo{tipo, tipo2},
	}
	db.Create(&servico)*/

	//Many to Many - Vários Serviços podem ser de Vários Tipos
	var tipos []Tipo
	err = db.Model(&Tipo{}).Preload("Servicos").Find(&tipos).Error
	if err != nil {
		panic(err)
	}

	//Exibindo todos os produtos de todas as categorias
	for _, tipo := range tipos {
		fmt.Println("Tipo:", tipo.ID, tipo.Nome, ":")
		for _, servico := range tipo.Servicos {
			fmt.Println("Servico:", servico.ID, servico.Nome, "Preço:", servico.Preco)
		}
	}
}
