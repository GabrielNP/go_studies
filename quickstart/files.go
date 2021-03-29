package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

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

func readFromFile() {
	fmt.Print("Reading with os.Open() method: ")
	file, err := os.Open("../sites-monitor/data/sites.txt")
	if err != nil {
		fmt.Println("An error ocurred:", err)
	}
	fmt.Println(file)
	fmt.Println("It returned the pointer to variable")
	fmt.Println()

	fmt.Print("Reading with ioutil.ReadFile() method:")
	f, e := ioutil.ReadFile("../sites-monitor/data/sites.txt")
	if e != nil {
		fmt.Println("Ocorreu um erro:", e)
	}
	fmt.Println()
	fmt.Println(string(f))
	fmt.Println()

	fmt.Print("Reading with bufio lib")
	reader := bufio.NewReader(file)
	for {
		line, er := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if er == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

func writeFile() {
	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		file.Close()
	}

	file.WriteString("Hello, world!")

	file.Close()
}
