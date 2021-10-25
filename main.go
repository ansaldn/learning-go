package main

import (
	"fmt"
	"runtime"
)

func main() {
	// welcome := "Hello World again"
	// var welcome1 string = "hellow world once more"
	// println(welcome)
	// println(welcome1)

	myAge := 21

	// var myAge int = 27

	// println("I am " + myAge + " years old")

	switch {
	case myAge >= 21:
		println("You are allowed to drink")
	case myAge <= 21 && myAge >= 18:
		println("You are almost of age to drink")
	default:
		println("You are too young to drink")

	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println("MacOS")
	case "linux":
		println("Linux Based Machine")
	default:
		println("Windows")

	}

	cities := [5]string{"London", "Manchester", "Leeds", "Birmingham", "Manchester"}
	fmt.Println(cities)

	var ukCities [5]string
	ukCities[2] = "Manchester"
	fmt.Println(ukCities)

}
