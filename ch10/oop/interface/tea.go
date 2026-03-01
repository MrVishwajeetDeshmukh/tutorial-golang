package main

import "fmt"

type Tea struct {
	name     string
	price    int
	category string
	taste    Taste
	state    State
}

func NewTea(
	name string,
	price int,
	category string,
	taste Taste,
	state State) *Tea {
	tea := new(Tea)
	tea.name = name
	tea.price = price
	tea.category = category
	tea.taste = taste
	tea.state = state
	return tea
}

// Make the tea
func (t *Tea) Make() error {
	if t.state == MakeDone {
		return fmt.Errorf("%s tea has already been made.", t.name)
	} else if t.state == Done {
		return fmt.Errorf(
			"%s tea has already been made and served to the customer.", t.name)
	}
	t.state = MakeDone
	return nil
}

// Package the tea
func (t *Tea) Package() error {
	if t.state == Waiting {
		return fmt.Errorf("%s tea is not ready yet.", t.name)
	} else if t.state == PackageDone {
		return fmt.Errorf("%s tea has already been packaged.", t.name)
	} else if t.state == Done {
		return fmt.Errorf(
			"%s tea has already been served to the customer.", t.name)
	}
	t.state = PackageDone
	return nil
}

// Serve the tea
func (t *Tea) Pick() error {
	if t.state == Waiting {
		return fmt.Errorf("%s tea is not ready yet.", t.name)
	} else if t.state == MakeDone {
		return fmt.Errorf("%s tea has not been packaged yet.", t.name)
	} else if t.state == Done {
		return fmt.Errorf(
			"%s tea has already been served to the customer.", t.name)
	}
	t.state = Done
	return nil
}

func (t *Tea) Name() string {
	return t.name
}

func (t *Tea) Price() int {
	return t.price
}

func (t *Tea) Category() string {
	return t.category
}

func (t *Tea) Taste() Taste {
	return t.taste
}

func (t *Tea) State() State {
	return t.state
}
