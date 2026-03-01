package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID   int
	Name string
}

var (
	tasks       []Task
	taskCounter int
	mutex       sync.Mutex
	wg          sync.WaitGroup
	taskChan    chan Task
)

func init() {
	// Log file configuration
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	log.SetOutput(file)
	log.Println("Program started")
}

func main() {
	taskChan = make(chan Task)
	go writeTasksToFile(taskChan)

	// Load to-do list from file on program start
	loadTasksFromFile()

	for {
		fmt.Println("To-Do Management Program")
		fmt.Println("1. Add a To-Do")
		fmt.Println("2. View To-Do List")
		fmt.Println("3. Exit")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			displayTasks()
		case 3:
			close(taskChan)
			wg.Wait()
			fmt.Println("Exiting program.")
			log.Println("Program exited")
			return
		default:
			fmt.Println("Please enter a valid number.")
		}
	}
}

func addTask() {
	fmt.Print("Enter To-Do name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove whitespace

	mutex.Lock()
	taskCounter++
	task := Task{ID: taskCounter, Name: name}
	tasks = append(tasks, task)
	mutex.Unlock()

	log.Printf("To-Do added: %v\n", task)

	wg.Add(1)
	taskChan <- task
}

func displayTasks() {
	mutex.Lock()
	defer mutex.Unlock()

	if len(tasks) == 0 {
		fmt.Println("No To-Dos registered.")
		return
	}

	fmt.Println("Current To-Do List:")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Name)
	}
}

func writeTasksToFile(ch <-chan Task) {
	file, err := os.OpenFile("tasks.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	for task := range ch {
		_, err := file.WriteString(fmt.Sprintf("%d,%s\n", task.ID, task.Name))
		if err != nil {
			log.Println("Failed to write to file:", err)
		} else {
			log.Printf("Saved To-Do to file: %v\n", task)
		}
		wg.Done()
	}
}

func loadTasksFromFile() {
	file, err := os.Open("tasks.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Existing To-Do file not found.")
			return
		}
		log.Println("Failed to read file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ",", 2)
		if len(parts) != 2 {
			continue
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		name := parts[1]
		task := Task{ID: id, Name: name}
		tasks = append(tasks, task)
		if id > taskCounter {
			taskCounter = id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error occurred while scanning file:", err)
	}

	log.Printf("Successfully loaded %d To-Dos.\n", len(tasks))
}
