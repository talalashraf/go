package main

import (
	"fmt"

	datastructure "github.com/talalashraf/go/pkg/datastructure/db"
)

func main() {
	// Setup
	fmt.Printf("Declare\n=======\n")
	d := &datastructure.Dimensions{}
	nv := &datastructure.NameValue{}

	// Retrieve
	fmt.Printf("Retrieving\n==========\n")
	d.Retrieve(10)
	nv.Retrieve("Talal")

	// Store
	fmt.Printf("Storing\n=======\n")
	d.X = 5
	d.Y = 15
	nv.Name = "Shameer"
	nv.Value = 7
	d.Store()
	nv.Store()

	// Sanity
	fmt.Printf("PrintStructs\n============\n%+v\n%+v\n", d, nv)
	fmt.Printf("PrintNameValue\n==============\n%+v\n%+v\n", nv.Name, nv.Value)
	fmt.Printf("PrintDimensions\n===============\n%+v\n%+v\n", d.X, d.Y)
	return
}
