package ttt

type Grid [][]string

type Options struct {
	TokTop    string
	ColorTop    string
	TokBtm    string
	ColorBtm    string
	Empty     string
	Dimension int
}
type Point struct {
	X int
	Y int
}
