package main

import "fmt"

func main() {
	type Student struct {
		Name       string
		Age        int
		Department string
		Score      map[string]int
	}

	student := Student{"Kim Chul-soo", 20, "Computer Engineering", map[string]int{
		"Science": 92,
		"Math":    94,
		"English": 95,
	}}

	fmt.Printf("struct student: %v\n", student)
	fmt.Println("================================================")
	fmt.Printf("name: %s\n", student.Name)
	fmt.Printf("age: %d\n", student.Age)
	fmt.Printf("department: %s\n", student.Department)
	fmt.Println("================================================")
	fmt.Println("[Subject Scores]")
	for name, score := range student.Score {
		fmt.Printf("%s: %d\n", name, score)
	}
	fmt.Println("================================================")
}
