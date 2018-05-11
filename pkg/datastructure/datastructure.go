package datastructure

// NameValue store name and value. Value is len name
type NameValue struct {
	Name     string
	Value    int64
	Store    func() error       // tbd database package
	Retrieve func(string) error // tbd database package
}

// Dimensions store X and Y. Y is X plus 10
type Dimensions struct {
	X        int64
	Y        int64
	Store    func() error      // tbd database package
	Retrieve func(int64) error // tbd database package
}
