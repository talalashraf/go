// Package db doesn't actually store in a database but proves the concept
// and reasoning for structuring code this way. Imagine blunt is your db wrapper
package db

import (
	"fmt"

	"github.com/talalashraf/go/pkg/datastructure"
)

// NameValue implements namevalue
type NameValue struct {
	datastructure.NameValue
	Talal int64
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
