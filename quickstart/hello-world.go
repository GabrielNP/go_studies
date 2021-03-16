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
			fmt.Println("Monitoring...")
		} else if option == 2 {
			fmt.Println("Showing logs...")
		} else if option == 3 {
			districts := showDistrictsWithArrays()
			fmt.Println(districts)
		} else if option == 0 {
			fmt.Println("Quiting...")
		} else {
			fmt.Println("Invalid option")
		}

		// Switch statement
		switch option {
		case 1:
			fmt.Println("Monitoring...")
			break
		case 2:
			fmt.Println("Showing logs...")
		case 3:
			districts := showDistrictsWithSlices()
			fmt.Println(districts)
		case 0:
			fmt.Println("Quiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid!")
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

func readChosenOption() int {
	// User input
	var option int
	fmt.Scanf("%d", &option)
	fmt.Println("The chosen command was:", option)

	fmt.Println("Please, confirm: ")
	fmt.Scan(&option) // Since we've already said option variable must be some integer value, with Scan() function we don't need to say the modifier
	fmt.Println()

	return option
}

func showMenu() {
	fmt.Println()
	fmt.Println("Please, choose an option:")
	fmt.Println("	1- Start Monitoring")
	fmt.Println("	2- Show Logs")
	fmt.Println("	3- Show Districts")
	fmt.Println("	0- Exit program")
}

func returnUserInfo() (string, int, string) {
	//  Declare variables
	var name string = "Gab"
	var age int = 24

	// Short variable declaration operator: another way to declare variable
	country := "Brazil"
	fmt.Println("country is a variable too:", country)

	return name, age, country
}

func showDistrictsWithArrays() [4]string {
	// Go lang arrays
	fmt.Println("Showing districts with arrays")

	var districts [4]string
	districts[0] = "RJ"
	districts[1] = "SP"
	districts[2] = "MG"
	districts[3] = "ES"

	return districts
}

func showDistrictsWithSlices() []string {
	// Go lang arrays
	fmt.Println("Showing districts with slices")

	districts := []string{"RJ", "SP", "MG", "ES"}

	return districts
}
