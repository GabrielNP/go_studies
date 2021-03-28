package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	pingDb()

	http.HandleFunc("/", index)
	fmt.Println("Listening server at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func connectDatabase() *sql.DB {
	connection := "host= user=postgres password= dbname=web_store sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDatabase()
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

	templates.ExecuteTemplate(w, "Index", produtcs)
}

func pingDb() {
	db := connectDatabase()
	defer db.Close()
}
