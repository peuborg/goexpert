package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	//Iniciando uma transação
	tx := db.Begin()
	var t Tipo
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&t, 1).Error
	if err != nil {
		panic(err)
	}

	//Commit na transação
	t.Nome = "Funilaria"
	tx.Debug().Save(&t)
	tx.Commit()
}
