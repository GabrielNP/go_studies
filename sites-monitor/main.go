package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	sayHello()

	for {
		showMenu()
		option := readChosenOption()

		switch option {
		case 1:
			startMonitoring()
			break
		case 2:
			fmt.Println("Showing logs...")
		case 0:
			fmt.Println("Quiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid!")
			os.Exit(-1)
		}
	}
}

func readChosenOption() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("O comando escolhido foi:", option)

	return option
}

func sayHello() {
	var name string = "Gab"
	var age int = 24
	var version float32 = 1.16

	fmt.Println("Hello World,", name, ",", age, "years old")
	fmt.Println("Runing Go on", version, "version")
}

func showMenu() {
	fmt.Println()
	fmt.Println("Please, choose an option:")
	fmt.Println("	1- Start Monitoring")
	fmt.Println("	2- Show Logs")
	fmt.Println("	0- Exit program")
}

func startMonitoring() {
	fmt.Println("\nMonitoring...")
	sites := []string{"https://random-status-code.herokuapp.com/"}

	for _, site := range sites {
		resp, _ := http.Get(site)
		if resp.StatusCode == 200 {
			fmt.Println("Site", site, "was successfully loaded")
		} else {
			fmt.Println("Site", site, "couldn't be loaded.\nStatus Code:", resp.Status)
		}

	}
}
