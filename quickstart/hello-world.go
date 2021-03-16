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
		} else if option == 0 {
			fmt.Println("Quiting program...")
			fmt.Println("Thanks for using...")
		} else {
			fmt.Println("Invalid option!")
		}

		// Switch statement
		switch option {
		case 1:
			fmt.Println(returnUserInfo())
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
	fmt.Println("	1- User info")
	fmt.Println("	2- Arrays")
	fmt.Println("	3- Slices")
	fmt.Println("	4- Loop Statements: basic for")
	fmt.Println("	5- Loop Statements: range for")
	fmt.Println("	6- Loop Statements: indeterminated for")
	fmt.Println("	7- map")
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
	// Arrays
	fmt.Println("This is a defined array with 4 elements")

	var districts [4]string
	fmt.Println("type:", reflect.TypeOf(districts))
	fmt.Println("length:", len(districts))
	fmt.Println("capacity:", cap(districts))
	fmt.Println()
	fmt.Println("Cannot has its length improved, i.e., doesn't accept more elements than those were previously defined")

	districts[0] = "RJ"
	districts[1] = "SP"
	districts[2] = "MG"
	districts[3] = "ES"

	return districts
}

func showDistrictsWithSlices() []string {
	// Slices
	fmt.Println("This is a slice")

	districts := []string{"RJ", "SP", "MG", "ES"}
	fmt.Println("type:", reflect.TypeOf(districts))
	fmt.Println("length:", len(districts))
	fmt.Println("capacity:", cap(districts))
	fmt.Println()

	fmt.Println("Adding one more element")
	districts = append(districts, "AM")

	fmt.Println("length:", len(districts))
	fmt.Println("capacity:", cap(districts))
	fmt.Println("Note that the inital capacity were doubled")
	fmt.Println()

	return districts
}

func showBasicForLoop() {
	districts := []string{"RJ", "SP", "MG", "ES"}

	for i := 0; i < len(districts); i++ {
		fmt.Println(districts[i])
	}
}

func showRangeForLoop() {
	districts := []string{"RJ", "SP", "MG", "ES"}

	for i, district := range districts {
		fmt.Println("Estou passando na posição", i, "do meu slice e essa posição tem o site", district)
	}
}

func showIndeterminatedFor() {
	i := 0
	for {
		if i > 10 {
			os.Exit(0)
		}
		fmt.Println(i)
		i++
	}
}

func showMap() {
	object := make(map[string]string)

	object["name"] = "Gabriel"
	object["age"] = "24"

	fmt.Println(object)
}
