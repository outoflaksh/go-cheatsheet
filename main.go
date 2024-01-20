package main

import "fmt"

func main() {
	// Printing on console
	fmt.Println("Hello, world!")

	// Variables
	var isEligible bool // Declaration without initializing
	var age int = 34
	var name string = "Lakshya"
	jobTitle := "Engineer" // Inferenced type

	// Conditional
	if age >= 18 {
		isEligible = true
	} else {
		isEligible = false
	}

	fmt.Printf("Eligibilty status of %v who is a %v = %v\n", name, jobTitle, isEligible)

	// Arrays
	// Fixed-length, indexed, homogenous type, contigous mem.
	var fruits [3]string = [3]string{"apple", "mango", "orange"}
	grades := [3]int32{98, 65, 87} // type inferred
	grades2 := [...]int32{100, 98} // length inferred

	fmt.Println(fruits[0], grades[0], grades2[0])

	// Slices
	// Wrappers around arrays; additional features

	var movies []string = []string{"Star Wars", "Gone Girl", "Spider-Man"} // initialised with no length value

	movies = append(movies, "Avengers")

	movies2 := []string{"Barbie", "Oppenheimer"}
	movies = append(movies, movies2...) // Spread operator

	fmt.Println(movies)

	fmt.Printf("Array is of length %v and capacity %v\n", len(movies), cap(movies))

	// Maps

	// Declaration
	var users map[string]int32 = make(map[string]int32)
	fmt.Println(users)

	// Initialize
	var users2 = map[string]int32{"Lakshya": 21, "John": 23}
	fmt.Println(users2["Lakshya"]) // Prints 21
	fmt.Println(users2["ABC"])     // Prints 0 in case of non existent key

	var _, doesExist = users2["James"] // Second option returns if the key exists
	fmt.Println(doesExist)

}
