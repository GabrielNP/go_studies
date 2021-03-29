package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	sayHello()

	// Golang does not have while statement. Passing for loop without params is become similar to it.
	for {
		showMenu()
		option := readChosenOption()

		// If statements
		if option == 1 {
			fmt.Println("Getting user info...")
		} else if option == 2 {
			fmt.Println("Arrays")
		} else if option == 3 {
			fmt.Println("Slices")
		} else if option == 4 {
			fmt.Println("Basic For Loops")
		} else if option == 5 {
			fmt.Println("Range For Loops")
		} else if option == 6 {
			fmt.Println("Indeterminated For Loops")
		} else if option == 7 {
			fmt.Println("Running map")
		} else if option == 8 {
			fmt.Println("Reading files")
		} else if option == 9 {
			fmt.Println("Writing files")
		} else if option == 10 {
			fmt.Println("Go Routine")
		} else if option == 0 {
			fmt.Println("Quiting program...\nThanks for using...")
		} else {
			fmt.Println("Invalid option!")
		}

		// Switch statement
		switch option {
		case 1:
			name, age, country := returnUserInfo()
			fmt.Printf("Name: %s\nAge: %d\nCountry: %s", name, age, country)
			break
		case 2:
			districts := showDistrictsWithArrays()
			fmt.Println(districts)
		case 3:
			districts := showDistrictsWithSlices()
			fmt.Println(districts)
		case 4:
			showBasicForLoop()
		case 5:
			showRangeForLoop()
		case 6:
			showIndeterminatedFor()
		case 7:
			showMap()
		case 8:
			readFromFile()
		case 9:
			writeFile()
		case 10:
			run()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Exiting with error...")
			os.Exit(-1)
		}
	}
}

func sayHello() {
	name, age, _ := returnUserInfo()
	var version float32 = 1.16

	// Print on terminal
	fmt.Println("Hello World,", name, ",", age, "years old")
	fmt.Println("Runing Go on", version, "version")

	// Infer types
	var city = "Sao Paulo"
	var simpleFloat = 1.6 // Go always choose float64 inferring variable types

	fmt.Println("Type of city variable is", reflect.TypeOf(city))
	fmt.Println("Type of simpleFloat variable is", reflect.TypeOf(simpleFloat))
}

func showMenu() {
	fmt.Println()
	fmt.Println("Please, choose an option:")
	fmt.Println("	1- User info")
	fmt.Println("	2- Arrays")
	fmt.Println("	3- Slices")
	fmt.Println("	4- Loop Statements: basic for")
	fmt.Println("	5- Loop Statements: range for")
	fmt.Println("	6- Loop Statements: indeterminated for")
	fmt.Println("	7- Map")
	fmt.Println("	8- Read File")
	fmt.Println("	9- Write File")
	fmt.Println("	10- Go routine")
	fmt.Println("	0- Exit program")
}

func returnUserInfo() (string, int, string) {
	//  Declare variables
	var name string = "Gab"
	var age int = 24

	// Short variable declaration operator: another way to declare variable
	country := "Brazil"

	return name, age, country
}
