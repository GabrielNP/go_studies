package main

import (
	"fmt"
	"os"
)

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
