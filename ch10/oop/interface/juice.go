package main

import "fmt"

type Juice struct {
	name     string
	price    int
	category string
	taste    Taste
	state    State
}

func NewJuice(
	name string,
	price int,
	category string,
	taste Taste,
	state State) *Juice {
	juice := new(Juice)
	juice.name = name
	juice.price = price
	juice.category = category
	juice.taste = taste
	juice.state = state
	return juice
}

// Make the juice
func (j *Juice) Make() error {
	if j.state == MakeDone {
		return fmt.Errorf("%s juice is already made.", j.name)
	} else if j.state == Done {
		return fmt.Errorf("%s juice has already been made and served to the customer.", j.name)
	}
	j.state = MakeDone
	return nil
}

// Package the juice
func (j *Juice) Package() error {
	if j.state == Waiting {
		return fmt.Errorf("%s juice is not ready yet.", j.name)
	} else if j.state == PackageDone {
		return fmt.Errorf("%s juice is already packaged.", j.name)
	} else if j.state == Done {
		return fmt.Errorf("%s juice has already been served to the customer.", j.name)
	}
	j.state = PackageDone
	return nil
}

// Serve the juice
func (j *Juice) Pick() error {
	if j.state == Waiting {
		return fmt.Errorf("%s juice is not ready yet.", j.name)
	} else if j.state == MakeDone {
		return fmt.Errorf("%s juice is not packaged yet.", j.name)
	} else if j.state == Done {
		return fmt.Errorf("%s juice has already been served to the customer.", j.name)
	}
	j.state = Done
	return nil
}

func (j *Juice) Name() string {
	return j.name
}

func (j *Juice) Price() int {
	return j.price
}

func (j *Juice) Category() string {
	return j.category
}

func (j *Juice) Taste() Taste {
	return j.taste
}

func (j *Juice) State() State {
	return j.state
}
