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

// Student represents a person who is a student with a specific grade level
type Student struct {
	Person // Embedding Person struct
	Grade  int
}

func main() {
	student := Student{Person: Person{Name: "Alice"}, Grade: 10}
	student.Greet()
	fmt.Println("I am in grade", student.Grade)
}
