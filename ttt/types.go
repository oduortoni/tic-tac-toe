package ttt

type Grid [][]string

type Options struct {
	SymbolTop    string
	ColorTop    string
	SymbolBtm    string
	ColorBtm    string
	Empty     string
	Dimension int
	Starter string
	Computer bool
	SymbolComp string
}
type Point struct {
	X int
	Y int
}
