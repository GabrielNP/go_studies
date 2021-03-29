package models

import (
	"go_studies/web-store-app/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SelectProducts() []Product {
	db := db.ConnectDatabase()
	defer db.Close()

	selectDb, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	produtcs := []Product{}

	for selectDb.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectDb.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		produtcs = append(produtcs, p)
	}

	return produtcs
}

func CreateProduct(name, description string, price float64, amount int) {
	db := db.ConnectDatabase()
	defer db.Close()

	insertDb, err := db.Prepare("INSERT INTO products (name, description, price, amount) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDb.Exec(name, description, price, amount)
}
