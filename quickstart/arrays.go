package main

import (
	"fmt"
	"reflect"
)

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
