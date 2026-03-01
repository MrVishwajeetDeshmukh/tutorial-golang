package main

import "fmt"

type Person struct {
	Name string
}

// Walk method for Person
func (p Person) Walk() {
	fmt.Println("I am walking.")
}

// Talk method for Person
func (p Person) Talk() {
	fmt.Printf("Hello, my name is %s.\n", p.Name)
}

type Student struct {
	Person
}

// Study method for Student
func (s Student) Study() {
	fmt.Println("I am studying.")
}

type Engineer struct {
	Person
}

// Develop method for Engineer
func (e Engineer) Develop() {
	fmt.Println("I am developing.")
}

type Reporter struct {
	Person
}

// Overridden Talk method for Reporter
func (r Reporter) Talk() {
	fmt.Printf("Hello, I am reporter %s from the news department.\n", r.Name)
}

// Publish method for Reporter
func (r Reporter) Publish(news string) {
	fmt.Printf("There is a latest report titled: %s\n", news)
}

func main() {
	// Creating different types of people
	person := Person{Name: "John Smith"}
	student := Student{Person: Person{Name: "Alex Kim"}}
	engineer := Engineer{Person: Person{Name: "Michael Johnson"}}
	reporter := Reporter{Person: Person{Name: "David Parker"}}

	person.Talk()
	student.Talk()
	engineer.Talk()
	reporter.Talk()
}
