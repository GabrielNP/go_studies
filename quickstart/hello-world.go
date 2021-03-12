package main

import (
	"fmt"
	"reflect"
)

func main() {
	//  Declare variables
	var name string = "Gab"
	var age int = 24
	var version float32 = 1.16

	// Print on terminal
	fmt.Println("Hello World,", name, ",", age, "years old")
	fmt.Println("Runing Go on", version, "version")

	// Another way to declare variable
	country := "Brazil"
	fmt.Println("country is a variable too:", country)

	// Infer types
	var city = "Sao Paulo"
	var simpleFloat = 1.6 // Go always choose float64 inferring variable types

	fmt.Println("Type of city variable is", reflect.TypeOf(city))
	fmt.Println("Type of simpleFloat variable is", reflect.TypeOf(simpleFloat))

	// User input
	fmt.Println()
	fmt.Println("Please, choose an option:")
	fmt.Println("	1- Iniciar Monitoramento")
	fmt.Println("	2- Exibir Logs")
	fmt.Println("	0- Sair do Programa")

	var option int
	fmt.Scanf("%d", &option)
	fmt.Println("O comando escolhido foi:", option)

	fmt.Println()
	fmt.Println("Please, confirm: ")

	fmt.Scan(&option) // Since we've already said option variable must be some integer value, with Scan() function we don't need to say the modifier
}
