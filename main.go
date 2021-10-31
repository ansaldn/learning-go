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
	// short form for arrays but each eliment has to be defined for anything to print
	ukCities[2] = "Manchester"
	fmt.Println(ukCities)

	citiesSlice := []string{"London", "Manchester", "Leeds", "Birmingham", "Manchester"}
	fmt.Println(citiesSlice)

	// Difference between arrays and slice is that the array has a predefined amount whereas the slice does not.

	var citiesVerbose []string
	citiesVerbose = append(citiesVerbose, "London")
	fmt.Println(citiesVerbose)


	ages := []int{10, 32, 89, 21, 45, 50, 17}

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("Hello World",i)
	}


	switch {
	case myAge >= 21:
		fmt.Println("Go Home")
	default:
		fmt.Println("Stay")
	
	fmt.Println(("Testing the waters with new change to directory"))
		
	}
		
	// Calling empy function
	HelloWorld()
	//Calling function with parameters passed through
	simpleIntro(21, "David")
	//Calling function with parameters passed through and output
	simpleMaths(9, 10)

	//Calling a struct back in main
	// When calling subsruct, use Struct: subScruct {subStruct variable: type}
	engineer := Engineer{
		Name: "David",
		Age: 21,
		Project: Project{
			Name: "Learning Go",
			projectID: 31223133,
			Technologies: []string{"Go"},
		},


	
	}
	fmt.Println(engineer)

	// Calling an interface
	feEnineer := &feEnineer{Name: "David"}
	PrintDetails(feEnineer)

}

// First function with no arguments passed through 
func HelloWorld()  {

	fmt.Println("Hello World")
	
}

// Functions with arguments passed through 
func simpleIntro(age int, name string)  {
	fmt.Println("My name is" , name)
	fmt.Println("I am", age, "years old")

	
}

// Function with arguments and expected output defined by "return"
func simpleMaths(a, b int) (int) {
	return a * a + b
	
}

//Struct 
type Engineer struct {
	Name string
	Age int
	Project Project
	
}
// "substruct as defined by 'Project Projecr' above "
type Project struct {
	Name string
	projectID int
	Priority string
	Technologies []string
	
}

func (e Engineer) amendAge() {
	e.Age += 1
	
}

// Interfaces in Go
type Employee interface {
	GetName() string
}

type feEnineer struct {

	Name string
	
}

func (e *feEnineer) GetName() string {
	return "Front End Eningeer Name: " + e.Name
}

func PrintDetails (e Employee) {
	fmt.Println(e.GetName())
}
