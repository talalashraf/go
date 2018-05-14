package ruth

// Data is our struct
type Data struct {
	X string
	Y int
	Z bool
}

// DataService is an interface that any db adapter must implement
type DataService interface {
	Store(*Data) error
	RetrieveByX(string) (*Data, error)
	RetrieveByY(int) (*Data, error)
	UpdateZ(*Data, bool) error
}
