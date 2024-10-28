package main

import "fmt"

// Person represents a person with a name
type Person struct {
	Name string
}

// Greet method for the Person struct
func (p *Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// Student represents a person who is a student with a specific college year
type Student struct {
	Person          // Embedding Person struct
	CollegeYear int // College year of the student
}

func main() {
	// Create a new Person and prompt for the user's name
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // Reads the user input and stores it in the 'name' variable

	// Create a new Student and prompt for the user's college year
	var year int
	fmt.Print("Enter your college year: ")
	fmt.Scanln(&year) // Reads the user input and stores it in the 'year' variable

	// Initialize a Student struct with the inputted name and college year
	student := Student{Person: Person{Name: name}, CollegeYear: year}

	// Greet the user and display their college year
	student.Greet()
	fmt.Println("I am in year", student.CollegeYear, "of college.")
}
