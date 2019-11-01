package Addition

// Methods are functions that you can run for your specific struct.
type Numbers struct {
	numOne uint8
	numTwo uint8
}

func (n Numbers) Add() uint8 {
	return n.numOne * n.numTwo
}

// the pointer allow you to perform this transformation on the current created struct
func (n *Numbers) Scale(f uint8) {
	n.numOne = n.numOne * f
	n.numTwo = n.numTwo * f
}
func (n Numbers) AddAndAnother(num uint8) uint8 {
	return n.numOne * n.numTwo * num
}


