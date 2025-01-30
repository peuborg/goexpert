package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pessoa struct {
	ID         int `gorm:"primaryKey;autoIncrement:true"`
	Nome       string
	Email      string
	gorm.Model //Cria as colunas datetime de ins upd del
}

func main() {
	//Open
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
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
	/*pessoas := []Pessoa{
		{Nome: "Bel", Email: "bel@email.com"},
		{Nome: "Cicl", Email: "cicl@email.com"},
		{Nome: "ze", Email: "ze@email.com"},
	}
	db.Create(pessoas)*/

	//Select one
	/*var pessoa1 Pessoa
	var pessoa2 Pessoa
	//Seleciona o primeiro registro com o código do parâmetro
	db.First(&pessoa1, 1)
	fmt.Println(pessoa1)
	//Seleciona o primeiro registro que satisfaz a condição
	db.First(&pessoa2, "nome = ?", "ze")
	fmt.Println(pessoa2)*/

	//Select all
	/*var pessoas []Pessoa
	db.Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}*/

	//Select os 5 primeiros registros
	/*var pessoas []Pessoa
	db.Limit(5).Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}*/

	//Select os 5 primeiros registros da página 2
	/*var pessoas []Pessoa
	db.Limit(5).Offset(5).Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}*/

	//WHERE
	/*var pessoas []Pessoa
	db.Where("id > ?", 5).Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}
	db.Where("nome LIKE ?", "%Cicl%").Find(&pessoas)
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}*/

	//UPDATE
	/*var p Pessoa
	db.First(&p, 1)
	p.Nome = "Beltraninho de Tal"
	db.Save(&p)
	//Seleciona o primeiro registro com o código do parâmetro
	db.First(&p, 1)
	fmt.Println(p.Nome)*/

	//DELETE
	/*var p Pessoa
	db.First(&p, 1)
	db.Delete(&p)*/

}
