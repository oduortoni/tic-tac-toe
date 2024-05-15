package ttt

type Grid [][]string

type Options struct {
	Token1    string
	Color1    string
	Token2    string
	Color2    string
	Empty     string
	Dimension int
}
type Point struct {
	X int
	Y int
}
