package edhi

// Dimensions store X and Y. Y is X plus 10
type Dimensions struct {
	X        int64
	Y        int64
	Store    func() error      // tbd database package
	Retrieve func(int64) error // tbd database package
}
