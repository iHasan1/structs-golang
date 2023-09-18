package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Anderson"} // Option 1 for declaring structs
	// alex := person{firstName: "Alex", lastName: "Anderson"} // Option 2 of declaring structs
	var alex person // Option 3 of declaring structs

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}