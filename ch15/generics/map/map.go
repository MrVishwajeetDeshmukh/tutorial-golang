package main

import "fmt"

// Define a generic map type
type Map[K comparable, V any] struct {
	data map[K]V
}

// Create a map
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V)}
}

// Add a key-value pair
func (m *Map[K, V]) Set(key K, value V) {
	m.data[key] = value
}

// Output value
func (m *Map[K, V]) Get(key K) (V, bool) {
	value, exists := m.data[key]
	return value, exists
}

func main() {
	// Map with string keys and integer values
	stringIntMap := NewMap[string, int]()
	stringIntMap.Set("a", 1)
	stringIntMap.Set("b", 2)
	fmt.Println(stringIntMap.Get("a")) // Output: 1 true
	fmt.Println(stringIntMap.Get("b")) // Output: 2 true

	// Map with integer keys and string values
	intStringMap := NewMap[int, string]()
	intStringMap.Set(1, "hello")
	intStringMap.Set(2, "world")
	fmt.Println(intStringMap.Get(1)) // Output: hello true
	fmt.Println(intStringMap.Get(2)) // Output: world true
}
