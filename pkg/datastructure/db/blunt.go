// Package blunt doesn't actually store in a database but proves the concept
// and reasoning for structuring code this way. Imagine blunt is your db wrapper
package db

import (
	"fmt"

	"github.com/talalashraf/go/pkg/datastructure"
)

// Dimensions implements dimensions
type Dimensions struct {
	datastructure.Dimensions
	Talal int64
}

// NameValue implements namevalue
type NameValue struct {
	datastructure.NameValue
	Talal int64
}

// Store will store the dimensions
func (d *Dimensions) Store() error {
	fmt.Printf("Storing dimension %d:%d\n", d.X, d.Y)
	return nil
}

// Retrieve will retrieve the dimensions
func (d *Dimensions) Retrieve(x int64) error {
	d.X = x
	d.Y = x + 10
	fmt.Printf("Retrieving dimension %d:%d\n", d.X, d.Y)
	return nil
}

// Store will store the dimensions
func (nv *NameValue) Store() error {
	fmt.Printf("Storing name:value %s:%d\n", nv.Name, nv.Value)
	return nil
}

// Retrieve will retrieve the dimensions
func (nv *NameValue) Retrieve(name string) error {
	nv.Name = name
	nv.Value = int64(len(name))
	fmt.Printf("Retrieving namevalue %s:%d\n", nv.Name, nv.Value)
	return nil
}
