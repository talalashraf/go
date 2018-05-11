package datastructure

// NameValue store name and value. Value is len name
type NameValue struct {
	Name     string
	Value    int64
	Store    func() error       // tbd database package
	Retrieve func(string) error // tbd database package
}
