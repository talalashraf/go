package main

import (
	"fmt"

	"github.com/talalashraf/go/pkg/ruth"
	db "github.com/talalashraf/go/pkg/ruth/mydb"
)

func main() {
	fmt.Printf("Jeeway\n")

	// Initiate our Ruth DB and create Interface
	c := db.Init("Talal's DB")
	ds := &db.DataService{Client: c}
	d := &ruth.Data{X: "Talal", Y: 123, Z: false}

	// Store
	e := ds.Store(d)
	if e != nil {
		fmt.Printf("Error: %v\n", e)
	}

	// Retrieve by X
	d, e = ds.RetrieveByX("Talal")
	if e != nil {
		fmt.Printf("Error: %v\n", e)
	}
	fmt.Printf("%+v\n", d)

	// Retrieve by Y
	d, e = ds.RetrieveByY(10)
	if e != nil {
		fmt.Printf("Error: %v\n", e)
	}
	fmt.Printf("%+v\n", d)

	// Store Z
	e = ds.UpdateZ(d, false)
	if e != nil {
		fmt.Printf("Error: %v\n", e)
	}
	fmt.Printf("%+v\n", d)

	fmt.Print("Bye\n")
	return
}
