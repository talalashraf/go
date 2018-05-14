package mydb

import (
	"fmt"

	"github.com/talalashraf/go/pkg/ruth"
)

// Ensure the interface is implemented
// Reference: https://splice.com/blog/golang-verify-type-implements-interface-compile-time/
// Reference: https://stackoverflow.com/a/13194635
var _ ruth.DataService = &DataService{}

// DataService implements DataService
type DataService struct {
	Client *Client
}

// Store mocks a store
func (ds *DataService) Store(d *ruth.Data) error {
	return ds.Client.Do(fmt.Sprintf("Store %+v", d))
}

// RetrieveByX mocks fetching Data by X
func (ds *DataService) RetrieveByX(x string) (*ruth.Data, error) {
	d := &ruth.Data{}
	d.X = x
	if e := ds.Client.Do(fmt.Sprintf("RetrieveByX %+v", d)); e != nil {
		return d, e
	}
	d.Y = 999
	d.Z = true
	return d, nil
}

// RetrieveByY mocks fetching by Y
func (ds *DataService) RetrieveByY(y int) (*ruth.Data, error) {
	d := &ruth.Data{}
	d.Y = y
	if e := ds.Client.Do(fmt.Sprintf("RetrieveByY %+v", d)); e != nil {
		return d, e
	}
	d.X = "ValueX"
	d.Z = true
	return d, nil
}

// UpdateZ mocks updating the bool
func (ds *DataService) UpdateZ(d *ruth.Data, z bool) error {
	d.Z = z
	return ds.Client.Do(fmt.Sprintf("UpdateZ %+v", d))
}
