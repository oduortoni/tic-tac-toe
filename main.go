package main

import (
	"bufio"
	"os"

	t "ttt/ttt"
)

func main() {
	reader := bufio.NewReader(os.Stdout)

	options := t.Options{
		SymbolTop:  "O",
		ColorTop:   t.GetColor("cyan"),
		SymbolBtm:  "X",
		ColorBtm:   t.GetColor("red"),
		Empty:      ".",
		Dimension:  3,
		Computer:   true,
		Starter:    "X",
		SymbolComp: "O",
	}

	t.Configure(reader, &options)

	t.Play(&options)
}
