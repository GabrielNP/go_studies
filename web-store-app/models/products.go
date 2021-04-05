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

	selectDb, err := db.Query("SELECT * FROM products ORDER BY id ASC")
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

		p.Id = id
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

func DeleteProduct(id string) {
	db := db.ConnectDatabase()
	defer db.Close()

	deleteDb, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteDb.Exec(id)
}

func EditProduct(id string) Product {
	db := db.ConnectDatabase()
	defer db.Close()

	editDb, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for editDb.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = editDb.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
	}

	return product
}

func UpdateProduct(name, description string, price float64, id, amount int) {
	db := db.ConnectDatabase()
	defer db.Close()

	updateDb, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, amount=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateDb.Exec(name, description, price, amount, id)
}
