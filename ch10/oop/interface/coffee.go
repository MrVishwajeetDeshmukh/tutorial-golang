package main

import "fmt"

type Coffee struct {
	name     string
	price    int
	category string
	taste    Taste
	state    State
}

func NewCoffee(
	name string,
	price int,
	category string,
	taste Taste,
	state State) *Coffee {
	coffee := new(Coffee)
	coffee.name = name
	coffee.price = price
	coffee.category = category
	coffee.taste = taste

	// added
	coffee.state = state
	return coffee
}

// Make the coffee
func (c *Coffee) Make() error {
	if c.state == MakeDone {
		return fmt.Errorf("%s coffee is already made.", c.name)
	} else if c.state == Done {
		return fmt.Errorf("%s coffee has already been made and served to the customer.", c.name)
	}
	c.state = MakeDone
	return nil
}

// Package the coffee
func (c *Coffee) Package() error {
	if c.state == Waiting {
		return fmt.Errorf("%s coffee is not ready yet.", c.name)
	} else if c.state == PackageDone {
		return fmt.Errorf("%s coffee is already packaged.", c.name)
	} else if c.state == Done {
		return fmt.Errorf("%s coffee has already been served to the customer.", c.name)
	}
	c.state = PackageDone
	return nil
}

// Serve the coffee
func (c *Coffee) Pick() error {
	if c.state == Waiting {
		return fmt.Errorf("%s coffee is not ready yet.", c.name)
	} else if c.state == MakeDone {
		return fmt.Errorf("%s coffee is not packaged yet.", c.name)
	} else if c.state == Done {
		return fmt.Errorf("%s coffee has already been served to the customer.", c.name)
	}
	c.state = Done
	return nil
}

func (c *Coffee) Name() string {
	return c.name
}

func (c *Coffee) Price() int {
	return c.price
}

func (c *Coffee) Category() string {
	return c.category
}

func (c *Coffee) Taste() Taste {
	return c.taste
}

func (c *Coffee) State() State {
	return c.state
}
