// Package db doesn't actually store in a database but proves the concept
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
